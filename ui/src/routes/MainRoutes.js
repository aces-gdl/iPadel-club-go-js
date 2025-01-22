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
const ImageUploadPage = Loadable(lazy(() => import('../pages/ImageUploadPage')));
const RankingPage = Loadable(lazy(() => import('../pages/Ranking')));


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
            path: '/upload-image',  
            element:
                <PrivateRoute>
                    <ImageUploadPage />
                </PrivateRoute>
        }
    ]
};

export default MainRoutes;
