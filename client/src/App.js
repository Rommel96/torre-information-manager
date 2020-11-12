import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Home from "./components/Home";
import Signup from "./components/Signup";
import Login from "./components/Login";
import NavBar from "./components/Navbar";

import { UserProvider } from "./context/UserContext";
import FindUser from "./components/FindUser";
import Favorites from "./components/Favorites";

function App() {
  return (
    <UserProvider>
      <Router>
        <NavBar />
        <div className="container p-4">
          <Switch>
            <Route path="/signup" component={Signup} />
            <Route path="/login" component={Login} />
            <Route path="/find-user" component={FindUser} />
            <Route path="/favorites" component={Favorites} />
            <Route path="/" component={Home} />
          </Switch>
        </div>
      </Router>
    </UserProvider>
  );
}

export default App;
