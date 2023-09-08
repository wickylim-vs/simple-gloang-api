# Simple GoLang API Application
You may use this simple GoLang API application, to practice for eLeisure 2.0 application deployment.

### Initializing & Running The Application Locally
1. Create a MySQL database instance (Tested with MySQL 8.1.0)
2. Initialize the MySQL database with the following details
    - MYSQL_DATABASE: api
    - MYSQL_USER: api
    - MYSQL_PASSWORD: password123
3. Run the GoLang application, supplying the database details using environment variables
    ```shell
    DB_HOST=localhost DB_PORT=13306 DB_USER=api DB_PASSWORD=password123 DB_NAME=api PORT=8080 go run app.go
    ```
4. Test the API
    ```shell
    curl http://localhost:8080/api/users
    ```

### What you can do to practice for eLeisure 2.0 application deployment
1. Fork this repository, or create a new GitHub repository for this GoLang app
2. Make the main branch a “protected branch”, so that
    - You will use Pull Request for code changes
3. Request for new Aliyun resources:
    - 1 ECS instance (to be used as both Jumphost & action runner)
    - 1 RDS instance, for MySQL database
4. Dockerize app, and keep the application Docker image in Aliyun ACR
5. Deploy the containerized app into Kubernetes (Aliyun ACK eLeisure Dev cluster)
6. Setup self-hosted GitHub actions runner (on the 1 ECS instance from step #3)
7. (After steps #1 to #5 is successful) Create GitHub Actions pipeline to automate deployment process
8. Explore the application platform on Aliyun
    - Monitor the ACK cluster & application (on Aliyun ACK dashboard or Aliyun Cloud Monitor)
    - Check application logs (using kubectl commands)

