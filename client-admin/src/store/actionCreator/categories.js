import {
  SET_CATEGORIES_SUCCESS,
  CREATE_CATEGORY_SUCCESS,
  DELETE_CATEGORY_SUCCESS,
  CATEGORY_LOADING,
  CATEGORY_ERROR,
} from '../actionTypes/category';

export const setCategories = () => async (dispatch) => {
  dispatch({
    type: CATEGORY_LOADING,
    payload: true,
  });
  try {
    const response = await fetch(`${process.env.REACT_APP_BASE_URL}/categories`);
    if (response.ok) {
      const result = await response.json();
      dispatch({
        type: SET_CATEGORIES_SUCCESS,
        payload: result,
      });
    } else {
      throw new Error(`HTTP error! status: ${response.status}`);
    };
  } catch (error) {
    dispatch({
      type: CATEGORY_ERROR,
      payload: error,
    });
  };
}

export const createCategory = (payload) => async (dispatch) => {
  dispatch({
    type: CATEGORY_LOADING,
    payload: true,
  });
  try {
    const response = await fetch(`${process.env.REACT_APP_BASE_URL}/categories`, {
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
        type: CREATE_CATEGORY_SUCCESS,
        payload: result,
      });
    } else {
      throw new Error(`HTTP error! status: ${response.status}`);
    };
  } catch (error) {
    dispatch({
      type: CATEGORY_ERROR,
      payload: error,
    });
  }
}

export const deleteCategory = (payload) => async (dispatch) => {
  dispatch({
    type: CATEGORY_LOADING,
    payload: true,
  });
  try {
    const response = await fetch(`${process.env.REACT_APP_BASE_URL}/categories/${payload}`, {
      method: 'DELETE',
      headers: {
        access_token: localStorage.getItem('access_token'),
      },
    });
    if (response.ok) {
      const result = await response.json();
      if (result) {
        dispatch({
          type: DELETE_CATEGORY_SUCCESS,
          payload,
        });
      };
    } else {
      throw new Error(`HTTP error! status: ${response.status}`);
    };
  } catch (error) {
    dispatch({
      type: CATEGORY_ERROR,
      payload: error,
    });
  }
}
