import createHistory from 'history/createHashHistory'
import * as React from 'react';
import { Route, Router, Switch } from 'react-router-dom';

import Profile from 'script/views/profile'
import Result from 'script/views/result';

import 'style/App.css';

const history = createHistory();


class App extends React.Component {
  public render() {
    return (
        <Router history={history}>
            <Switch>
                <Route exact={true} path="/" component={Result}/>
                <Route path="/profile" component={Profile}/>
            </Switch>
        </Router>
    );
  }
}

export default App;
