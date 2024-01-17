import grpc from 'k6/experimental/grpc';
import { check } from 'k6';
import { generateRandomValidMailAddress, generateRandomEmailAddressWithValidDomain, generateInvalidMailAdress } from './utils/mail.js';

let runName = __ENV.RUN_NAME || "generic_run_" + Math.floor(Math.random() * (9999999 - 1000000) + 1000000);
let mailVerifierHost = __ENV.MAIL_VERIFIER_HOST || "localhost"
let mailVerifierPort = __ENV.MAIL_VERIFIER_PORT || "9090"
let mailVerifierAddress = `${mailVerifierHost}:${mailVerifierPort}`

const client = new grpc.Client();
client.load(['../../app/protos/'], 'service.proto');

export let options = {
  scenarios: {
    syntaxVerification: {
      executor: 'ramping-vus',
      exec: 'syntaxVerification',
      startVUs: 0,
      tags: {
        runName: runName,
      },
      stages: [
        { duration: '10s', target: 100 },
        { duration: '30s', target: 65 },
        { duration: '5s', target: 200 },
        { duration: '10s', target: 0 },
      ],
      gracefulRampDown: '5s',
    },
    simpleVerification: {
      executor: 'ramping-vus',
      exec: 'simpleVerification',
      startVUs: 0,
      tags: {
        runName: runName,
      },
      stages: [
        { duration: '10s', target: 70 },
        { duration: '35s', target: 50 },
        { duration: '10s', target: 0 },
      ],
      gracefulRampDown: '5s',
    },
    fullVerification: {
      executor: 'ramping-vus',
      exec: 'fullVerification',
      startVUs: 0,
      tags: {
        runName: runName,
      },
      stages: [
        { duration: '10s', target: 30 },
        { duration: '30s', target: 5 },
        { duration: '5s', target: 80 },
        { duration: '10s', target: 0 },
      ],
      gracefulRampDown: '5s',
    },
  }
};

export function syntaxVerification() {
  if (__ITER === 0) {
    client.connect(mailVerifierAddress, {
      plaintext: true
    });
  }

  let syntaxVerification = client.invoke('MailVerifier/SyntaxVerification', {
    email: getRandomEmail(generateRandomValidMailAddress, generateInvalidMailAdress)
  });

  check(syntaxVerification, {
    'success': (r) => r && r.status === grpc.StatusOK,

  });
}

export function simpleVerification() {
  if (__ITER === 0) {
    client.connect(mailVerifierAddress, {
      plaintext: true
    });
  }

  let simpleVerification = client.invoke('MailVerifier/SimpleVerification', {
    email: getRandomEmail(generateRandomEmailAddressWithValidDomain, generateInvalidMailAdress)
  });

  check(simpleVerification, {
    'success': (r) => r && r.status === grpc.StatusOK,
  });
}

export function fullVerification() {
  if (__ITER === 0) {
    client.connect(mailVerifierAddress, {
      plaintext: true
    });
  }

  let fullVerification = client.invoke('MailVerifier/FullVerification', {
    email: getRandomEmail(generateRandomEmailAddressWithValidDomain, generateInvalidMailAdress)
  });

  check(fullVerification, {
    'success': (r) => r && r.status === grpc.StatusOK,
  });
}

export function teardown() {
  client.close();
}

function getRandomEmail(validMailGenerator, invalidMailGenerator) {
  if (Math.random() > 0.30) {
    return validMailGenerator();
  } else {
    return invalidMailGenerator();
  }
}