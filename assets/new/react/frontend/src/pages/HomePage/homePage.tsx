import * as React from 'react';

import { Example } from './components/Example';
import { FetchNote } from './components/FetchNote';

export class HomePage extends React.Component {
  public render() {
    return (
      <div>
        <h1 className='page-title'>Home</h1>
        <Example text='new text' />
        <FetchNote />
      </div>
    );
  }
}
