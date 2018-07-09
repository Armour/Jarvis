import { EXAMPLE_ACTION } from 'constants/actions';
import { IActionExampleAction, IExample } from 'types';

const initialState: IExample = {
  text: 'initial text',
};

export const example = (state = initialState, action: IActionExampleAction): IExample => {
  switch (action.type) {
    case EXAMPLE_ACTION:
      return {
        text: state.text + ', ' + action.text,
      };
    default:
      return state;
  }
};
