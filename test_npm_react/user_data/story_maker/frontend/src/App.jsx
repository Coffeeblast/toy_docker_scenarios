import { useState, useReducer} from 'react'
import './App.css'
import constants from "./constants.js"
import {getStories} from "./backend_communication.js"
import {
  storiesReducer,
  storiesContext,
  dispatchStoriesContext,
} from "./stories_reducer.js"
import {ScreenStoriesList} from "./ScreenStoriesList.jsx"
import { ScreenStoriesButtons } from './ScreenStoriesButtons.jsx'
import { AddStoryForm } from './AddStoryForm.jsx'

function App() {
  // state variables
  const [stories, dispatchStories] = useReducer(storiesReducer, null);
  const [addStoryFormData, setAddStoryFormData] = useState(constants.makeAddStoryFormDataEmpty())
  const [addStoryActive, setAddStoryActive] = useState(false);
  // state variables end

  async function refreshStories(){
    // as reducer itself cannot be asynchronous function, we must do the asynchronous part outside of it BEFORE dispatching an action
    try{
      const jsonData = await getStories(); 
      let oldStoriesMap = new Map() // ket is story.id, value is story.selected
      if (stories !== null) {
        stories.forEach((story) => oldStoriesMap.set(story.id, story.selected))
      }
      const newStories = (jsonData === null) ? [] : jsonData.map(
        (jd) => {
          return {...jd, selected: oldStoriesMap.has(jd.id) ? oldStoriesMap.get(jd.id): false}}
        );
      dispatchStories({type: "refresh_all", data: newStories});
    }
    catch(error) {
      console.error("Failed to refresh stories:", error);
      dispatchStories({type: "refresh_all", data: []});
    }
  }

  const storiesFetched = (stories !== null);
  if (!storiesFetched){
    refreshStories();
  };

  const selectedStoriesCount = (stories === null) ? 0 : stories.reduce(
    (soFar, val) => soFar + ((val.selected) ? 1 : 0),
    0,
  )

  let addStory = <></>;
  if (addStoryActive){
    addStory = (
      <div style={{"padding" : "10px"}}>
        <AddStoryForm 
          formData={addStoryFormData} 
          setFormData={setAddStoryFormData}
          setAddStoryActive={setAddStoryActive}
          refreshStories={refreshStories}
        />
      </div>
    )
  }
  return (
    <storiesContext.Provider value={stories}>
      <dispatchStoriesContext.Provider value={dispatchStories}>
        <div style={{"display" : "flex"}}>
          <div style={{"padding" : "10px"}}>
            <ScreenStoriesList storiesFetched={storiesFetched}/>
          </div>
          <div style={{"padding" : "10px"}}>
            <ScreenStoriesButtons 
              setAddStoryFormData={setAddStoryFormData} 
              setAddStoryActive={setAddStoryActive} 
              addStoryActive={addStoryActive}
              oneStorySelected={selectedStoriesCount === 1}
            />
          </div>
          {addStory}
        </div>
      </dispatchStoriesContext.Provider>
    </storiesContext.Provider>
  ) ; 
}

export default App
