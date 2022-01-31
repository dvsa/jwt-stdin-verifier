```
       _          _                      _  __
      (_)_      _| |_    __   _____ _ __(_)/ _|_   _
      | \ \ /\ / / __|___\ \ / / _ \ '__| | |_| | | |
      | |\ V  V /| ||_____\ V /  __/ |  | |  _| |_| |
     _/ | \_/\_/  \__|     \_/ \___|_|  |_|_|  \__, |
    |__/                                       |___/
    Verify JWT's as part of an Apache Rewrite rule
```
## What does it do

This program will listen to stdin and treat any input as a JWT. The JWT will be verified and either `OK` printed or 
`INVALID - {error}` for a valid or invalid token. It is designed to be used in conjuction with Apache Rewrite to allow 
access based on a valid JWT within a query string


## Usage

Pretty simple: `./jwt-verify -keyfile <path/to/keyfile>`

| Argumnet | Value                      | Required |
|----------|----------------------------|----------|
| keyfile  | Path to a RS256 public key | Yes      | 

### Keyfile

We make the assumption that the JWT to verify has been signed using an RS256 algorithm. Because of this you will need to
provide the public key in a PEM file format.

### Generating a key

The following commands can be used to create the required private and public key with the names `jwtRS256.key` & `jwtRS256.pub`

```
ssh-keygen -t rsa -P "" -b 4096 -m PEM -f jwtRS256.key
ssh-keygen -e -m PEM -f jwtRS256.key > jwtRS256.key.pub
```


## Apache integration

Below is an example how this can be used with Apache Rewrite rules

```
RewriteEngine On
  RewriteMap jwtprg "prg:/usr/local/bin/jwt-verify -keyfile /path/to/keyfile" apache:apache
  <Location "/foo">
    Require all granted
    LogLevel alert rewrite:trace8
    Options FollowSymLinks

    # Set JWT_AUTH to false first
    RewriteRule ^ - [ENV=JWT_AUTH:false]

    # Set JWT_AUTH environment variable if jwt-verify has authenticated the request
    RewriteCond %{QUERY_STRING} ^jwt=([^&]+)
    RewriteCond ${jwtprg:%1} ^OK$
    RewriteRule ^ - [ENV=JWT_AUTH:true]

    # Disallow access if JWT_AUTH environment variable is not true
    RewriteCond %{ENV:JWT_AUTH} !^true$
    RewriteRule ^ - [F]

    ErrorDocument 403 "Please <a href='/login'>authenticate</a>."

  </Location>
```
