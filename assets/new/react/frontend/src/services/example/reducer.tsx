import { EXAMPLE_ACTION } from './constants';
import { IActionExampleAction, IExampleState } from './types';

const initialState: IExampleState = {
  text: 'initial text',
};

export const example = (state = initialState, action: IActionExampleAction): IExampleState => {
  switch (action.type) {
    case EXAMPLE_ACTION:
      return {
        text: state.text + ', ' + action.text,
      };
    default:
      return state;
  }
};
