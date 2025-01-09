from pydantic import BaseModel

class Story(BaseModel):
    id: int
    name: str
    author: str