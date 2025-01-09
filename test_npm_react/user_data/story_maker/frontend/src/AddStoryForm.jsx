import {postAddStory} from "./backend_communication.js"

export function AddStoryForm({formData, setFormData, setAddStoryActive, refreshStories}){
    const onFormChange = (evt) => handleAddStoryFormChange(evt, formData, setFormData)
    return (
        <form>
        <h4>New story:</h4>
        <label htmlFor="name">Name:</label> <br/>
        <input type="text" id="name" name="name" value={formData.name} onChange={onFormChange} />
        <br/>
        <label htmlFor="author">Author:</label> <br/>
        <input type="text" id="author" name="author" value={formData.author} onChange={onFormChange} /> <br/>
        <button type="button" onClick={() => handleAddStoryConfirm(setAddStoryActive, formData, refreshStories)}> OK </button>
        <button type="button" onClick={() => handleAddStoryCancel(setAddStoryActive)}> Cancel </button>
        </form>
    );
}

async function handleAddStoryConfirm(setAddStoryActive, formData, refreshStories){
    setAddStoryActive(false);
    try {
      const jsonData = await postAddStory(formData);
      if (jsonData.success)
      {
        alert("Story added successfully");
        refreshStories();
      }
      else {
        console.error('Add story failed.');
      }
    }
    catch (error) {
      console.error("Add story failed: ", error);
    }
  }

  function handleAddStoryCancel(setAddStoryActive){
    setAddStoryActive(false);
  }

  function handleAddStoryFormChange(evt, formData, setFormData) {
    const newData = {...formData, [evt.target.name] : evt.target.value};
    setFormData(newData);
  }