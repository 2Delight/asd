import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './Utils/reportWebVitals';
import Router from "./Router/Router";
import './Translation/Translation'

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
      <Router/>
  </React.StrictMode>
);

reportWebVitals();
