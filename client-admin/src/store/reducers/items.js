import {
  SET_ITEMS_SUCCESS,
  CREATE_ITEM_SUCCESS,
  UPDATE_ITEM_SUCCESS,
  DELETE_ITEM_SUCCESS,
  ITEM_LOADING,
  ITEM_ERROR,
} from '../actionTypes/item';

const initialState = {
  items: [],
  isLoading: false,
  error: null,
}

export default function items(state = initialState, action) {
  switch (action.type) {
    case SET_ITEMS_SUCCESS:
      return {
        ...state,
        items: action.payload,
        isLoading: false,
      };
    case CREATE_ITEM_SUCCESS:
      return {
        ...state,
        items: [
          action.payload,
          ...state.items,
        ],
        isLoading: false,
        error: null,
      };
    case UPDATE_ITEM_SUCCESS:
      return {
        ...state,
        items: [
          action.payload,
          ...state.items.filter(item => item.id !== action.id),
        ],
        isLoading: false,
        error: null,
      };
    case DELETE_ITEM_SUCCESS:
      return {
        ...state,
        items: state.items.filter(item => item.id !== action.payload),
        isLoading: false,
        error: null,
      };
    case ITEM_LOADING:
      return {
        ...state,
        isLoading: action.payload,
      };
    case ITEM_ERROR:
      return {
        ...state,
        error: action.payload,
        isLoading: false,
      };
    default:
      return state;
  }
}