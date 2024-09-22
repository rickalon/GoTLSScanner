# Go HTTPS TLS Scanner API
![image](https://github.com/user-attachments/assets/41ac172b-54d7-46d2-89cb-1f597f0e8be8)

# Description 
API that enables users to input an array of URLs. For each URL provided, the application performs an analysis of the associated TLS certificates. It retrieves and organizes detailed information about each certificate, including its subject's common name, issuer, expiration and issuance dates, public key algorithm, and any associated DNS names. This tool helps users monitor and assess the security of their URLs by providing insights into the certificates that secure their connections.

To enhance efficiency, a MongoDB storage solution is implemented to persist the URLs. This allows users to avoid repeated and expensive GET calls, improving overall performance. Additionally, a login and registration system using OAuth2 is integrated, ensuring secure access for users and managing their sessions effectively.

# Endpoints
## Submit URLs for Analysis
- **Endpoint**: `POST /api/v1/analyze`
- **Request Body**:
    ```json
    {
      "urls": [
        "https://www.google.com/",
        "https://www.example.com/"
      ]
    }
    ```
  - **Response**:
  ```json
    [
  {
    "url": "https://www.trivago.com",
    "urls": [
      {
        "url": "https://www.trivago.com",
        "resutl": "Found",
        "to": "imgio.trivago.com",
        "From": "DigiCert TLS RSA SHA256 2020 CA1",
        "country": [
          "US"
        ],
        "expDate": "2025-04-22 23:59:59 +0000 UTC",
        "emiDate": "2024-04-22 00:00:00 +0000 UTC",
        "alg": "ECDSA",
        "dns": [
          "imgio.trivago.com",
          "www.youzhan.com",
          "www.trivago.vn",
          "www.trivago.tw",
          "www.trivago.tv",
          "www.trivago.sk",
          "www.trivago.si",
          "www.trivago.sg",
          "www.trivago.se",
          "www.trivago.rs",
          "www.trivago.ro",
          "www.trivago.pt",
          "www.trivago.pl",
          "www.trivago.ph",
          "www.trivago.pe",
          "www.trivago.no",
          "www.trivago.nl",
          "www.trivago.net",
          "www.trivago.mx",
          "www.trivago.kr",
          "www.trivago.jp",
          "www.trivago.jobs",
          "www.trivago.it",
          "www.trivago.in",
          "www.trivago.ie",
          "www.trivago.hu",
          "www.trivago.hr",
          "www.trivago.hk",
          "www.trivago.gr",
          "www.trivago.fr",
          "www.trivago.fi",
          "www.trivago.es",
          "www.trivago.ec",
          "www.trivago.dk",
          "www.trivago.de",
          "www.trivago.cz",
          "www.trivago.com.vn",
          "www.trivago.com.uy",
          "www.trivago.com.tw",
          "www.trivago.com.tr",
          "www.trivago.com.sg",
          "www.trivago.com.pt",
          "www.trivago.com.ph",
          "www.trivago.com.pe",
          "www.trivago.com.my",
          "www.trivago.com.mx",
          "www.trivago.com.hr",
          "www.trivago.com.hk",
          "www.trivago.com.ec",
          "www.trivago.com.cy",
          "www.trivago.com.co",
          "www.trivago.com.br",
          "www.trivago.com.au",
          "www.trivago.com.ar",
          "www.trivago.com",
          "www.trivago.co.za",
          "www.trivago.co.uk",
          "www.trivago.co.th",
          "www.trivago.co.nz",
          "www.trivago.co.kr",
          "www.trivago.co.in",
          "www.trivago.co.il",
          "www.trivago.co.id",
          "www.trivago.cl",
          "www.trivago.ch",
          "www.trivago.cat",
          "www.trivago.ca",
          "www.trivago.bg",
          "www.trivago.be",
          "www.trivago.at",
          "www.trivago.ae",
          "jsa.youzhan.com",
          "jsa.trivago.com",
          "imgio.youzhan.com",
          "imgim.youzhan.com",
          "imgim.trivago.com",
          "imgcy.youzhan.com",
          "ar.trivago.com"
        ],
        "isCA": false
      },
      {
        "url": "https://www.trivago.com",
        "resutl": "Found",
        "to": "DigiCert TLS RSA SHA256 2020 CA1",
        "From": "DigiCert Global Root CA",
        "country": [
          "US"
        ],
        "expDate": "2031-04-13 23:59:59 +0000 UTC",
        "emiDate": "2021-04-14 00:00:00 +0000 UTC",
        "alg": "RSA",
        "dns": null,
        "isCA": true
      }
    ]
  },
  {
    "url": "https://www.google.com/",
    "urls": [
      {
        "url": "https://www.google.com/",
        "resutl": "Found",
        "to": "www.google.com",
        "From": "WR2",
        "country": [
          "US"
        ],
        "expDate": "2024-11-18 07:15:48 +0000 UTC",
        "emiDate": "2024-08-26 07:15:49 +0000 UTC",
        "alg": "ECDSA",
        "dns": [
          "www.google.com"
        ],
        "isCA": false
      },
      {
        "url": "https://www.google.com/",
        "resutl": "Found",
        "to": "WR2",
        "From": "GTS Root R1",
        "country": [
          "US"
        ],
        "expDate": "2029-02-20 14:00:00 +0000 UTC",
        "emiDate": "2023-12-13 09:00:00 +0000 UTC",
        "alg": "RSA",
        "dns": null,
        "isCA": true
      },
      {
        "url": "https://www.google.com/",
        "resutl": "Found",
        "to": "GTS Root R1",
        "From": "GlobalSign Root CA",
        "country": [
          "BE"
        ],
        "expDate": "2028-01-28 00:00:42 +0000 UTC",
        "emiDate": "2020-06-19 00:00:42 +0000 UTC",
        "alg": "RSA",
        "dns": null,
        "isCA": true
      }
    ]
  }
    ]
    ```
    
# Deployment
The deployment is compose of the golang image and a mongodb image
Build the Dockerfile with the tag tlsapi
Configure the enviroment variables that the api and mongo are going to use
 - DBUSER=root
 - DBPASS=example
 - DBPORT=27017
 - DBHOST=localhost
Run de dockercompose file.
You can deploy the application in any cloud provider.
