import { EXAMPLE_ACTION } from './constants';

// Example state
export interface IExampleState {
  text: string;
}

// Notes actions
export interface IActionExampleAction {
  type: typeof EXAMPLE_ACTION;
  text: string;
}

export type IActionsExample =
  | IActionExampleAction;
