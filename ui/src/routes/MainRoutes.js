import { lazy } from 'react';

// project imports
import MainLayout from '../layout/MainLayout';
import Loadable from './Loadable';
import { Navigate } from 'react-router-dom';


const PrivateRoute = ({ children }) => {
    const isAuthenticated = localStorage.getItem('user'); // Asume que guardas el user en localStorage
    return isAuthenticated ? children : <Navigate to="/auth/login" />;
};

const Home = Loadable(lazy(() => import('../pages/Home')));
const UpdatePasswordPage = Loadable(lazy(() => import('../pages/UpdatePasswordPage')));
const AboutUsPage = Loadable(lazy(() => import('../pages/AboutUsPage')));
const EventRegistration = Loadable(lazy(() => import('../pages/EventRegistration')));
const RankingPage = Loadable(lazy(() => import('../pages/Ranking')));
const Images = Loadable(lazy(() => import('../pages/images')));

// ==============================|| MAIN ROUTING ||============================== //

const MainRoutes = {
    path: '/',
    element: <MainLayout />,
    children: [
        {
            path: '/',
            element:
                <PrivateRoute>
                    <Home />
                </PrivateRoute>
        },
        {
            path: '/update-password',
            element:
                <PrivateRoute>
                    <UpdatePasswordPage />
                </PrivateRoute>
        },
        {
            path: '/aboutus',
            element:
                <PrivateRoute>
                    <AboutUsPage />
                </PrivateRoute>
        },
        {
            path: '/inscriptions',
            element:
                <PrivateRoute>
                    <EventRegistration />
                </PrivateRoute>
        },
        {
            path: '/ranking',
            element:
                <PrivateRoute>
                    <RankingPage />
                </PrivateRoute>
        },    
        {
            path: '/images',  
            element:
                <PrivateRoute>
                    <Images/>
                </PrivateRoute>
        }
    ]
};
export default MainRoutes;
