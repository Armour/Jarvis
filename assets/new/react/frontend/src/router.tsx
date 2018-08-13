import * as React from 'react';
import { Route, Switch } from 'react-router';

import 'materialize-css';
import 'sass/global';

import { Footer } from 'components/Footer';
import { Header } from 'components/Header';
import { HomePage } from 'pages/HomePage';
import { NotFoundPage } from 'pages/NotFoundPage';

export const router = (
  <div>
    <Header />
    <Switch>
      <Route exact path='/' component={HomePage} />
      <Route component={NotFoundPage} />
    </Switch>
    <Footer />
  </div>
);
