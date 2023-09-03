import {
  SET_CATEGORIES_SUCCESS,
  CREATE_CATEGORY_SUCCESS,
  DELETE_CATEGORY_SUCCESS,
  CATEGORY_LOADING,
  CATEGORY_ERROR,
} from '../actionTypes/category';

const initialState = {
  categories: [],
  isLoading: false,
  error: null,
}

export default function items(state = initialState, action) {
  switch (action.type) {
    case SET_CATEGORIES_SUCCESS:
      return {
        ...state,
        categories: action.payload,
        isLoading: false,
      };
    case CREATE_CATEGORY_SUCCESS:
      return {
        ...state,
        categories: [
          action.payload,
          ...state.categories,
        ],
        isLoading: false,
      };
    case DELETE_CATEGORY_SUCCESS:
      return {
        ...state,
        categories: state.categories.filter(category => category.id !== action.payload),
        isLoading: false,
      };
    case CATEGORY_LOADING:
      return {
        ...state,
        isLoading: action.payload,
      };
    case CATEGORY_ERROR:
      return {
        ...state,
        error: action.payload,
        isLoading: false,
      };
    default:
      return state;
  }
}