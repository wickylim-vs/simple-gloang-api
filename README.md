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
    go mod tidy
    go build -o app .
    DB_HOST=localhost DB_PORT=13306 DB_USER=api DB_PASSWORD=password123 DB_NAME=api PORT=8080 ./app
    ```
4. Test the API
    ```shell
    curl http://localhost:8080/api/users
    ```

### What you can do to practice for eLeisure 2.0 application deployment
1. Fork this repository, or create a new GitHub repository for this GoLang app (within "gentingmalaysia" organization)
2. Make sure "gentingmalaysia" organization's self-hosted runners can be used from the new repository
3. Make the main branch a “protected branch”, so that
    - You will use Pull Request for code changes
4. Get the Dockerfile for this GoLang app from the developers
    - Update the base images to use "golang:latest" in gentingmalaysia's ACR registry
5. Also, get the Kubernetes deployment yaml files for this GoLang app from the developers
    - Update the application and database images to use images in gentingmalaysia's ACR registry
6. Create GitHub Actions pipeline, by updating the provided template pipeline script (referencing Chowkidar's pipeline) 
7. Configure GitHub environment, variables, and secrets as required by the pipeline script
8. Make code changes to trigger the pipeline for deployment, make sure all the steps are successful
9. Once the app is deployed successfully, explore the application on Aliyun platform
    - Monitor the ACK cluster & application (on Aliyun ACK dashboard)
    - Check application logs (using kubectl logs commands)

