# Sovertis

## Configuring .env file for this project

In order to configure and run this project successfully, you need to set up a `.env` file to manage environment variables. The `.env` file will contain sensitive information and configuration settings needed for the project to work correctly.

Follow the steps below to create and configure your `.env` file:

### Step 1: create a new .env file

Create a new file in the root directory of the project.

   - For example in Linux terminal: 

        ```
        $ touch .env
        ```

### Step 2: define environment variables

In the `.env` file, define environment variables using the `KEY=VALUE` syntax.

   - For example:

     ```
     DB_HOST=postgres
     DB_PORT=5432
     SECRET_KEY=mysecretkey
     DEBUG=True
     ```

   - Replace the example values with your specific configuration.


### Step 3: run the project

With the `.env` file properly configured, you can now run this project.

   - Use the appropriate commands or scripts to start this project.

   - The environment variables defined in the `.env` file will be available for services defined in the `docker-compose.yml` file.
