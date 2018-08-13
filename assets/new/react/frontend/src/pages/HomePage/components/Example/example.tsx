import * as React from 'react';
import { connect } from 'react-redux';
import { AnyAction, Dispatch } from 'redux';

import { exampleAction } from 'services/example/actions';
import { IGlobalState } from 'types/global';

const styles = require('./example.scss');

// Component

interface IExampleComponentStateProps {
  text: string;
}

interface IExampleComponentDispatchProps {
  setText(): void;
}

type IExampleComponentProps = IExampleComponentStateProps & IExampleComponentDispatchProps;

export class ExampleComponent extends React.Component<IExampleComponentProps> {
  public onClick = (e: React.MouseEvent<HTMLAnchorElement>) => {
    e.preventDefault();
    this.props.setText();
  }

  public render() {
    return (
      <div>
        <div className='center-align'>
          <a href='#' className='btn todo-filter-btn waves-effect waves-light' onClick={this.onClick}>Append Text</a>
        </div>
        <div className={`card-panel center-align z-depth-2 ${styles['example-card-panel']}`}>
          <span className='teal-text'>
            {this.props.text}
          </span>
        </div>
      </div>
    );
  }
}

// Container

interface IExampleProps {
  text: string;
}

const mapStateToProps = (state: IGlobalState): IExampleComponentStateProps => ({
  text: state.exampleState.text,
});

const mapDispatchToProps = (dispatch: Dispatch<AnyAction>, ownProps: IExampleProps): IExampleComponentDispatchProps => ({
  setText: () => {
    dispatch(exampleAction(ownProps.text));
  },
});

export const Example = connect(
  mapStateToProps,
  mapDispatchToProps,
)(ExampleComponent);
