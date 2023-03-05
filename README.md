# Make a monetized API with Go, Google Cloud Run and RapidAPI

- Go - The language to code the server
- Google Cloud Run - Where the server is hosted
- RapidAPI - Where to handle the APIs

---

**NOTE!**  
The secret key that is sent from RapidAPI to GCP is X-Rapidapi-Proxy-Secret, but somehow it's key is X-Mashape-Proxy-Secret when it reaches GCP.

---

You need to add an .env file with the following key-value

`X_MASHAPE_PROXY_SECRET=1111...`

This key is found in RapidAPI under the name X-Rapidapi-Proxy-Secret.

