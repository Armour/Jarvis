import { combineReducers } from 'redux';

import { example } from 'services/example/reducer';
import { notes } from 'services/notes/reducer';
import { IGlobalState } from 'types/global';

export default combineReducers<IGlobalState>({
  exampleState: example,
  notesState: notes,
});
