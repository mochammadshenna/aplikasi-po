import {
  SET_ITEMS_SUCCESS,
  CREATE_ITEM_SUCCESS,
  UPDATE_ITEM_SUCCESS,
  DELETE_ITEM_SUCCESS,
  ITEM_LOADING,
  ITEM_ERROR,
} from '../actionTypes/item';

export const setItems = () => async (dispatch) => {
  dispatch({
    type: ITEM_LOADING,
    payload: true,
  });
  try {
    const response = await fetch(`${process.env.REACT_APP_BASE_URL}/items`);
    if (response.ok) {
      const result = await response.json();
      dispatch({
        type: SET_ITEMS_SUCCESS,
        payload: result,
      });
    } else {
      const error = await response.json();
      throw (error);
    };
  } catch (error) {
    dispatch({
      type: ITEM_ERROR,
      payload: error,
    });
  };
}

export const createItem = (payload) => async (dispatch) => {
  dispatch({
    type: ITEM_LOADING,
    payload: true,
  });
  try {
    const response = await fetch(`${process.env.REACT_APP_BASE_URL}/items`, {
      method: 'POST',
      headers: {
        "Content-Type": "application/json",
        access_token: localStorage.getItem('access_token'),
      },
      body: JSON.stringify(payload),
    });
    if (response.ok) {
      const result = await response.json();
      dispatch({
        type: CREATE_ITEM_SUCCESS,
        payload: result,
      });
    } else {
      const error = await response.json();
      throw (error);
    };
  } catch (error) {
    dispatch({
      type: ITEM_ERROR,
      payload: error,
    });
  }
}

export const updateItem = (id, payload) => async (dispatch) => {
  dispatch({
    type: ITEM_LOADING,
    payload: true,
  });
  try {
    const response = await fetch(`${process.env.REACT_APP_BASE_URL}/items/${id}`, {
      method: 'PUT',
      headers: {
        "Content-Type": "application/json",
        access_token: localStorage.getItem('access_token'),
      },
      body: JSON.stringify(payload),
    });
    if (response.ok) {
      const result = await response.json();
      dispatch({
        type: UPDATE_ITEM_SUCCESS,
        payload: result,
        id,
      });
    } else {
      const error = await response.json();
      throw (error);
    };
  } catch (error) {
    dispatch({
      type: ITEM_ERROR,
      payload: error,
    });
  }
}

export const deleteItem = (payload) => async (dispatch) => {
  dispatch({
    type: ITEM_LOADING,
    payload: true,
  });
  try {
    const response = await fetch(`${process.env.REACT_APP_BASE_URL}/items/${payload}`, {
      method: 'DELETE',
      headers: {
        access_token: localStorage.getItem('access_token'),
      },
    });
    if (response.ok) {
      const result = await response.json();
      if (result) {
        dispatch({
          type: DELETE_ITEM_SUCCESS,
          payload,
        });
      };
    } else {
      const error = await response.json();
      throw (error);
    };
  } catch (error) {
    dispatch({
      type: ITEM_ERROR,
      payload: error,
    });
  }
}
