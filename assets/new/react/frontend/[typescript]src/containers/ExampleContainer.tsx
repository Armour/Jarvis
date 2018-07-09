import { connect } from 'react-redux';
import { AnyAction, Dispatch } from 'redux';

import { exampleAction } from 'actions';
import { ExampleComponent, IExampleComponentDispatchProps, IExampleComponentStateProps } from 'components/ExampleComponent';
import { IAppState } from 'types';

export interface IExampleContainerProps {
  text: string;
}

const mapStateToProps = (state: IAppState): IExampleComponentStateProps => ({
  text: state.example.text,
});

const mapDispatchToProps = (dispatch: Dispatch<AnyAction>, ownProps: IExampleContainerProps): IExampleComponentDispatchProps => ({
  setText: () => {
    dispatch(exampleAction(ownProps.text));
  },
});

export const ExampleContainer = connect(
  mapStateToProps,
  mapDispatchToProps,
)(ExampleComponent);
