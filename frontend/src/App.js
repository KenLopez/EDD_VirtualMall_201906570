import React from 'react'
import {BrowserRouter as Router, Route, Redirect} from 'react-router-dom'
import Home from './components/Home'
import CargarTienda from './components/CargarTienda'
import CargarPedido from './components/CargarPedido'
import CargarProducto from './components/CargarProducto'
import Reporte from './components/Reporte'
import NavBar from './components/NavBar'

function App() {
  return (
    <>
     
      <Router>
        <NavBar/>
        <Route exact path="/">
          <Redirect to="/Home" />
        </Route>
        <Route path="/Home" component={Home}/>
        <Route path="/CargarTienda" component={CargarTienda}/>
        <Route path="/CargarPedido" component={CargarPedido}/>
        <Route path="/CargarProducto" component={CargarProducto}/>
        <Route path="/Reporte" component={Reporte}/>
      </Router>
    </>
  )
}

export default App
