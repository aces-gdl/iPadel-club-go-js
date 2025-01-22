import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import { ThemeProvider } from './themes/ThemeContext'; // Importar el Provider de tema
import { BrowserRouter } from 'react-router-dom';
const root = ReactDOM.createRoot(document.getElementById('root'));

const config = {
    // basename: only at build time to set, and Don't add '/' at end off BASENAME for breadcrumbs, also Don't put only '/' use blank('') instead,
    // like '/berry-material-react/react/default'
    // fontFamily
    presetColor: 'default',
    basename: '/',
    defaultPath: '/',
    fontFamily: `'Roboto', sans-serif`,
    borderRadius: 12
};

root.render(
    <BrowserRouter basename={config.basename}>
        <ThemeProvider>
            <App />
        </ThemeProvider>
    </BrowserRouter>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
