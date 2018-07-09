import { List, Map, Set } from 'immutable';

// Constants
import { EXAMPLE_ACTION, RECEIVE_NOTE } from 'constants/actions';
import { DEFAULT_RECEIVE_ERROR, RECEIVE_RESPONSE, START_REQUEST } from 'constants/actions';

// Components Interfaces
export { IExampleComponentStateProps, IExampleComponentDispatchProps } from 'components/ExampleComponent';
export { INoteFetcherStateProps, INoteFetcherDispatchProps } from 'components/FetchNote';
export { IExampleContainerProps } from 'containers/ExampleContainer';

// Global State
export interface IAppState {
  example: IExample;
  fetching: IFetchingSet;
  notes: INoteList;
}

// Example
export interface IExample {
  text: string;
}

// Fetching
export type IFetchingSet = Set<string>;

// Notes
export interface INote {
  id: number;
  content: string;
}
export type INoteList = List<INote>;

// Example Actions
export interface IActionExampleAction {
  type: typeof EXAMPLE_ACTION;
  text: string;
}

// Fetching Actions
export interface IActionStartRequest {
  type: typeof START_REQUEST;
  url: string;
  method: string;
}
export interface IActionReceiveResponse {
  type: typeof RECEIVE_RESPONSE;
  url: string;
  method: string;
}
export interface IActionDefaultReceiveError {
  type: typeof DEFAULT_RECEIVE_ERROR;
  url: string;
  error: string;
}
export type IActionsFetchApi =
  | IActionStartRequest
  | IActionReceiveResponse
  | IActionDefaultReceiveError

// Notes Actions
export interface IActionReceiveNote {
  type: typeof RECEIVE_NOTE;
  data: {
    notes: INoteList;
    code: number;
  };
}
export type IActionsFetchNote =
  | IActionReceiveNote

