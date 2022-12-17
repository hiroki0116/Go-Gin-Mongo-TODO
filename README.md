## TaskGO

This is super simple task manager app. You can keep track of your daily tasks easily. This was developed to learn Golang with Gin framework.

## Production Site 

*FYI, API server is deployed on Cloud Run and it will cold start. You might need to wait for 15 seconds to run the application.
[https://go-next-tasks.vercel.app/](https://go-next-tasks.vercel.app/)


<div style="flex:100%">
 <img src="https://res.cloudinary.com/sixty-seconds-idea-training-project/image/upload/v1671121225/ApplicationLayout/firstHalfTask_eizrqv.gif" width="45%"/>
 <img src="https://res.cloudinary.com/sixty-seconds-idea-training-project/image/upload/v1671154890/ApplicationLayout/shortClip_zbjpbt.gif" width="45%"/>
</div>

<img src="https://res.cloudinary.com/sixty-seconds-idea-training-project/image/upload/v1671155191/ApplicationLayout/shortClip_2_vgezfa.gif" width="45%">

## Getting Start

1. Install front end project on your local machine

Frontend repository [https://github.com/hiroki0116/go-nextjs-todo](https://github.com/hiroki0116/go-nextjs-todo)

2. Install npm
```bash
 $ npm install
``` 

3. Run local server
```bash
 $ npm run dev
```
4. Install API server on your local machine and run docker container
```bash
 $ docker compose up -d --build
```
