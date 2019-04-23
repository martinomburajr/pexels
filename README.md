# pexels
An application that pulls images from [Pexels][1] and adds them as your background photo. 

## 1. Introduction
The pexels application is written entirely in [Golang][2] and currently runs as a binary that acts as a HTTP server. This server accepts requests that perform actions such as: 

    1. Fetching New Photos: _Given a Pexel Image ID_
    2. Fetching Random Photos: 
    3. Automatically Setting a Downloaded Image as a backdrop.
    
## 2. Installation
The current way of installing the application is to use Docker. The image is in no current dockerhub repository (waiting for test coverage to be sufficient). 

### 2.1 Docker
  1. Download the repository either as a `zip` or via `git clone`. 
  2. Ensure you have [docker installed][3] on your machine.
  3. Build the image using the following command
  
  `docker build -t pexels:0.1-RC-1 .`
  
  4. Run once successfully built (check console for errors). Run the container with the following command:
  
    `docker run  -p 9191:9191 --restart=always  -v /home/<your-user>/.pexels:/home/appuser/.pexels  --name pexels-daemon pexels:0.1-RC-1`
    
    Ensure to change <your-user> to the actual user.
    
  5. The container should be running in the background.
  
#### 2.1.1 Things to note about Docker Install
   1. As the application is running in a separate linux environment (docker container). It currently cannot access the hosts system to change the background. This will have to be done manually in the `/home/<your-user>/.pexels/pictures` directory.
   2. By default the application runs on `port 9191`. If you want to change this, this currently can only be done in source and `Dockerfile`. An issue to modify this is in the pipeline
    
    
## 2. Using the application 
The application currently only runs as a server so some HTTP request will need to be made in order to interact with the server

**1. Get a Photo By ID**

 `curl localhost:9191/new/{id}`

Where {id} is an integer value id that corresponds to a pexels image.


**2. Get Random Photo**

 `curl localhost:9191/rand?size="original"`
 
This returns a random photo from their curated page. The size parameter is optional so you can exclude it. If excluded, it shall download their `large` size by default. The `original` is Pexel's highest quality image. A sizing guide can be seen by running the command in the next section.

**3. Get Photo Sizing Information**

 `curl localhost:9191/sizes`
 
This outputs their sizes that can be used as HTTP URL query parameters when wanting to obtain pictures of a certain size and quality. Once you have read the size description run any command that retrieves the photos with a given size parameter e.g. `curl localhost:9191/rand?size="portrait"`
  
  
## Contributions
Play around with the application and create some issues if you encounter a bug, or would like to see something added.

#### Disclaimer
I am not affiliated with Pexels, and this is not their Official Application. I just thought it would be a fun little tool.

[1]: https://pexels.com
[2]: https://golang.org
[3]: https://docs.docker.com/install/