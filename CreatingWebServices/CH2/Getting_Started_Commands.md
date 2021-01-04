# Getting Started List of Commands
## Environment Setup
### [6:12]
```
docker run --rm -it -d^
   -v %CD%/data:/var/lib/mysql ^
   -p 3306:3306 ^
   -e MYSQL_ROOT_PASSWORD=password123 ^
   mysql:latest
```