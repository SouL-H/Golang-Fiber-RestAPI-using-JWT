
# Golang-Fiber-RestAPI using JWT

### Starting (M1-ARM64)

```
go mod tidy

```

## .env File

>You need to create a .env config file.

```env
DB_CONN="root:*******@/dbName" #DB conf 
SECRET_KEY="test_secret_key" #JWT Secret Key
PORT=":8082" #Port to listen
```

#

```
Register

Post localhost:8082/api/register

```

<p>
    <img src="./img/register.png"  style="width:350px;" alt="Observer">

</p>

#

```
Login

Post localhost:8082/api/login

```

<p>
    <img src="./img/login.png"  style="width:350px;" alt="Observer">

</p>

#

```
Get UserInfo

```
> Cookie JWT is saved.
<p>
    <img src="./img/getUser.png"  style="width:350px;" alt="Observer">

</p>

#

```
Logout

Post localhost:8082/api/logout

```
> User information is deleted from the cookie.
<p>
    <img src="./img/logout.png"  style="width:350px;" alt="Observer">

</p>


# Cookie

<p>
    <img src="./img/cookie.png"  style="width:850px;" alt="Observer">

</p>


####  Thank You Scalable Scripts
