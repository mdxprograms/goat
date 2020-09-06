import React from 'react';
import { Route, Switch, BrowserRouter as Router } from 'react-router-dom'

const Home = () => <h1>hello goat</h1>

const App = () => (
  <Router>
    <Switch>
      {/* more routes here */}
      <Route exact path='/'>
        <Home />
      </Route>
    </Switch>
  </Router>
)

export default App;
