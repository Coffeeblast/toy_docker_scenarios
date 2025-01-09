const __BACKEND_URL = "http://localhost:8000/" // must end with /

const constants = {
    urlGetStories: `${__BACKEND_URL}stories`,
    urlPostAddStory: `${__BACKEND_URL}add-story`,
    // this is creation function, so that noone mutates the object
    makeAddStoryFormDataEmpty: (() => {return {"id" : 0, "name" : "", "author" : ""};}), 
};

export default constants;