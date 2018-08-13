import { EXAMPLE_ACTION } from './constants';
import { IActionExampleAction } from './types';

export const exampleAction = (text: string): IActionExampleAction => ({
  type: EXAMPLE_ACTION,
  text,
});
