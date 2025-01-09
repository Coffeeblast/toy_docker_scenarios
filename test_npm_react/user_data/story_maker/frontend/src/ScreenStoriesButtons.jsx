import constants from "./constants.js"

export function ScreenStoriesButtons({setAddStoryFormData, setAddStoryActive, addStoryActive, oneStorySelected}){
    const buttonStyle = {"display" : "block"}
    const addButtonDisabled = addStoryActive
    const editButtonDisabled = addStoryActive || (! oneStorySelected)
    return (
      <>
        <button 
          style={buttonStyle} 
          onClick={() => handleAddStoryStart(setAddStoryFormData, setAddStoryActive)}
          disabled={addButtonDisabled}
        > Add 
        </button>
        <button 
          style={buttonStyle}
          disabled={editButtonDisabled}
        > Edit 
        </button>
      </>
    )
}

function handleAddStoryStart(setAddStoryFormData, setAddStoryActive){
  setAddStoryFormData(constants.makeAddStoryFormDataEmpty());
  setAddStoryActive(true);
}