import React from 'react';
import { Auth0Provider } from '@auth0/auth0-react';
import { PropsWithChildren } from 'react';

const AuthProvider: React.FC<PropsWithChildren> = ({ children }) => {
  return (
    <Auth0Provider
      domain="dev-1ig767b0haje6gfw.us.auth0.com"
      clientId="VfIe5K5gwWf9KsUhOHEbEuME663XtaYg"
      authorizationParams={{
        redirect_uri: import.meta.env.VITE_AUTH0_CALLBACK_URL,
        audience: import.meta.env.VITE_AUTH0_AUDIENCE,
        scope: "openid profile email read:services write:services"
      }}
    >
      {children}
    </Auth0Provider>
  );
};

export default AuthProvider;