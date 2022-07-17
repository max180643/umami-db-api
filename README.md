# Umami-DB-API

API for Umami Database (Only supported for PostgreSQL)

</details>

### Create a shortened URL with custom alias

##### API request

`POST /track`

<details>
<summary>Body (JSON)</summary>

```json
{
    "websiteName": "[umami-website-name]", #required
    "hostname": "[umami-website-domain]", #required
    "ip": "[request-ip]",
    "userAgent": "[request-user-agent]",
    "url": "[request-path]"
}
```

</details>

##### API response

<details>
<summary>JSON</summary>

```json
{
  "status": "success",
  "response": "Save successfully."
}
```

</details>

</details>

### Health Check

##### API request

`GET /health`

##### API response

<details>
<summary>JSON</summary>

```json
{
  "Status": "success",
  "Response": "OK"
}
```

</details>

---

## Environment variables

create config.json file

```
{
    "SERVER": {
        "DEBUG": "", # gin.Mode (true/false)
        "PORT": "" # both "config.json" and "Dockerfile" need to configure on same port
    },
    "DB": {
        "ADDRESS": "",
        "PORT": "",
        "USER": "",
        "PASSWORD": "",
        "DBNAME": ""
    }
}
```

---

<sub><sup>
This project includes IP2Location LITE data available from [https://lite.ip2location.com](https://lite.ip2location.com)
</sup></sub>
