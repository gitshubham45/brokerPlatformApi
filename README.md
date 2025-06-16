##  Dependencies

### Local Machine Setup

1. **Go**: v1.20 or newer  
2. **Docker** and **Docker Compose**  
3. **API-Testing**: `curl` or `postman` for testing APIs  

---

##  Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/gitshubham45/brokerPlatformApi.git 
cd brokerPlatformApi
```

### 2. Set Up Environment Variables
- Create a .env file:

```bash
    cp .env.example .env
```

### 3. Start mongo DB in docker

```
docker run -d \
  --name mongodb \
  -p 27017:27017 \
  mongo
```

### 4. Start go server

```
go run main.go
```

### 5 Test the APIS

```
curl --location 'http://localhost:8080/api/user/signup' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "email" : "example@abc.com",
    "password" : "abc1234"
}'

-----------------------------------------------------------



```


