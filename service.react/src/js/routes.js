'use strict';

import React from 'react'; // eslint-disable-line no-unused-vars
import { Route } from 'react-router';
import Application from './components/Application';
import HomePage from './components/HomePage';
import UserPage from './components/UserPage';
import RepoListPage from './components/RepoListPage';
import RepoPage from './components/RepoPage';

export default (
  <Route component={Application}>
    <Route path="/" component={HomePage} />
    <Route path="/:username" component={UserPage} />
    <Route path="/:username/repos" component={RepoListPage} />
    <Route path="/:username/:reponame" component={RepoPage} />
  </Route>
);
