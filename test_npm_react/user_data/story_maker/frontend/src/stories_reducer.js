import { createContext } from 'react';

export function storiesReducer(stories /*state*/, action) { // returns a new state 
    if (action.type === "refresh_all") {
        return action.data;
    } else if (action.type === "story_selected_change") {
        const newStories = stories.map((story) => {return (action.id == story.id) ? {...story, selected: !story.selected} : story})
        return newStories
    } else {
        throw Error('Unknown action: ' + action.type);
    }
}

export const storiesContext = createContext(null);
export const dispatchStoriesContext = createContext(null);