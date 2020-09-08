import React, { useEffect, useState } from 'react';
import { Route, Switch, BrowserRouter as Router } from 'react-router-dom'

const Home = () => {
  const [ping, setPing] = useState("")

  const getPing = () => {
    fetch("/api/ping").then(response => response.json())
      .then(({ Message }) => setPing(Message))
  }

  useEffect(() => {
    getPing()
  }, [])

  return <h1>{ping}</h1>
}

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
