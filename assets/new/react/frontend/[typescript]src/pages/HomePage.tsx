import * as React from 'react';

import { ExampleContainer } from 'containers/ExampleContainer';
import { FetchNote } from 'containers/FetchNote';

export class HomePage extends React.Component {
  public componentDidMount() {
    // ...
  }

  public render() {
    return (
      <div>
        <h1>Home</h1>
        <ExampleContainer text='test' />
        <FetchNote />
      </div>
    );
  }
}
