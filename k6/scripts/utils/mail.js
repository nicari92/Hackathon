const DOMAIN_NAME_EXTENSIONS = ['com', 'it', 'io', 'net', 'org', 'gov', 'edu', 'biz', 'info']
const VALID_DOMAINS = ['gmail.com', 'hotmail.com', 'yahoo.com', 'outlook.com', 'icloud.com', 'aol.com', 'zoho.com', 'protonmail.com', 'gmx.com', 'yandex.com', 'mail.com']

export function generateRandomValidMailAddress() {
  return `${randomString()}@${randomString()}.${DOMAIN_NAME_EXTENSIONS[Math.floor(Math.random() * DOMAIN_NAME_EXTENSIONS.length)]}`;
}

export function generateRandomEmailAddressWithValidDomain() {
  return `${randomString()}@${VALID_DOMAINS[Math.floor(Math.random() * VALID_DOMAINS.length)]}`;
}

export function generateInvalidMailAdress() {
  return `${randomString()}${randomString()}`;

}

function randomString() {
  return Math.random().toString(36).substring(2, 7);
}