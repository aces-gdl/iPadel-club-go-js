import { lazy } from 'react';

// project imports
import Loadable from './Loadable';
import MinimalLayout from '../layout/MinimalLayout';

// login option 3 routing
const LoginPage = Loadable(lazy(() => import('../pages/LoginPage')));
const RegisterPage = Loadable(lazy(() => import('../pages/RegisterPage')));

// ==============================|| AUTHENTICATION ROUTING ||============================== //

const AuthenticationRoutes = {
    path: '/',
    element: <MinimalLayout />,
    children: [
        {
            path: '/auth/login',
            element: <LoginPage/>
        },
        {
            path: '/auth/register',
            element: <RegisterPage />
        }
    ]
};

export default AuthenticationRoutes;
