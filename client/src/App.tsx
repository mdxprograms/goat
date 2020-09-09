import React, { useEffect, useState } from 'react';
import { Route, Switch, BrowserRouter as Router } from 'react-router-dom'

type HomeStyle = {
  background?: string,
  color?: string,
  padding?: string
}

const homeStyle: HomeStyle = {
  background: "#444",
  color: "lightblue",
  padding: "12px",
}

const Home: React.FC = () => {
  const [ping, setPing] = useState("")

  const getPing = () => {
    fetch("/api/ping").then(response => response.json())
      .then(({ Message }) => setPing(Message))
  }

  useEffect(() => {
    getPing()
  }, [])

  return <h1 style={homeStyle}>{ping}</h1>
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
