import constants from "./constants.js"

export async function getStories(){
    console.log("Getting data for stories");
    const response = await fetch(
        constants.urlGetStories, 
        {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
            },
        }
    );
    const jsonData = await response.json();
    console.log("data returned: ", {jsonData});
    return jsonData;
}
  
export async function postAddStory(story){
    const response = await fetch(constants.urlPostAddStory, {
        method: 'POST', // HTTP method
        headers: {
        'Content-Type': 'application/json', // Specify JSON format
        },
        body: JSON.stringify(story), // Convert data to JSON string
    }
    );
    if (!response.ok) {
    throw new Error('Add story: Network response was not ok ' + response.statusText);
    }
    return response.json(); // Parse JSON response  
}