from pydantic import BaseModel, Field, field_validator
from typing import Annotated

class Book(BaseModel):
    id: int
    title: Annotated[str, Field(min_length=1)]
    # title: str = Field(min_length=1)
    author: str
    year: int
    isbn: str

    @field_validator("year")
    def validate_year(cls, value):
        if 1000 <= value <= 2026:
            return value
        raise ValueError("Year must be in the range 1000 and 2026")
    
    @field_validator("isbn")
    def validate_isbn(cls, value):
        if len(value) not in (10, 13):
            raise ValueError("Length of isbn must e 10 or 13")
        return value
    
class ResponseModel(BaseModel):
    message: str
    book: Book