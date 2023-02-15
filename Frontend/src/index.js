import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import { combineReducers, configureStore } from '@reduxjs/toolkit'
import { roleSlice } from './features/role/roleSlice';
import { Provider } from 'react-redux';
import { persistStore, persistReducer } from 'redux-persist';
import storage from 'redux-persist/lib/storage';
import { PersistGate } from 'redux-persist/integration/react';
import { getDefaultMiddleware } from '@reduxjs/toolkit';


const customizedMiddleware = getDefaultMiddleware({
  serializableCheck: false
})

const persistConfig = {
  key: 'root',
  storage,
};

const persistedReducer = persistReducer(persistConfig, roleSlice.reducer);


const store = configureStore({
  reducer: persistedReducer,
  middleware : customizedMiddleware
  
})
const persistor = persistStore(store);


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <Provider store={store}>
    <PersistGate loading={null} persistor={persistor}>
    <App />
    </PersistGate>
    </Provider>
  </React.StrictMode>
);

