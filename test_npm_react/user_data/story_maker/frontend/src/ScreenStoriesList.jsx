import {useContext} from 'react'
import {
    storiesContext,
    dispatchStoriesContext,
  } from "./stories_reducer.js"

function RowStory({story}){
    const dispatchStories = useContext(dispatchStoriesContext);
    return (
      <tr>
        <td>
          <input 
            type="checkbox" 
            checked={story.selected} 
            onChange={() => dispatchStories({type: "story_selected_change", id: story.id} )}
          />
        </td>
        <td>{story.author}</td>
        <td>{story.name}</td>
      </tr>
    );
  }
  
  function TableStories(){
    const stories = useContext(storiesContext)
    return (
      <table>
        <thead>
          <tr>
            <th />
            <th>author</th>
            <th>name</th>
          </tr>
        </thead>
        <tbody>
          {stories.map((story) => <RowStory key={story.id} story={story}/>)}
        </tbody>
      </table>
    );
  }
  
  export function ScreenStoriesList({storiesFetched}){
    const storyContent = (storiesFetched) ? (
      <TableStories/>
    ) : (<button> Loading data </button>);
    return (
      <>
        <h1>Stories</h1>
        {storyContent}
      </>
    );
  }