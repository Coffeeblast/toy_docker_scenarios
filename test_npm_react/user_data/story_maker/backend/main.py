from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
#import asyncio
import data_model as dm
from database import Database

DATABASE = Database()

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:5173"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@app.get("/")
async def root():
    return {"message": "Hello World"}

@app.get("/stories")
async def get_stories():
    #await asyncio.sleep(5) # simulate high network latency
    return DATABASE.get_stories()

@app.post("/add-story")
async def add_story(story: dm.Story):
    DATABASE.add_story(story)
    return {"success" : True}