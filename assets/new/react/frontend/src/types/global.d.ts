import { RouterState } from 'connected-react-router';

import { IExampleState } from 'services/example/types';
import { INotesState } from 'services/notes/types';

// Global state
export interface IGlobalState {
  notesState: INotesState;
  exampleState: IExampleState;
  router?: RouterState;
}

// Interface for async call steps
export interface IAsyncCall {
  REQUESTED: string;
  SUCCESS: string;
  FAILURE: string;
}
