import React, { useState, useEffect } from 'react';
import { ShoppingCart, Package, History, LogOut, User, Home, Plus, Minus, Trash2, Star, Phone, Mail, MapPin } from 'lucide-react';

const API_BASE_URL = 'http://localhost:8080';

const App = () => {
  const [currentScreen, setCurrentScreen] = useState('login');
  const [currentPage, setCurrentPage] = useState('home');
  const [user, setUser] = useState(null);
  const [token, setToken] = useState(localStorage.getItem('token'));
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(false);
  const [cart, setCart] = useState(null);
  const [orders, setOrders] = useState([]);

  // Sample product images - you can replace with your own or add image_url to your database
  const productImages = {
    1: 'https://images.unsplash.com/photo-1496181133206-80ce9b88a853?w=400&h=400&fit=crop', // Laptop
    2: 'https://images.unsplash.com/photo-1527864550417-7fd91fc51a46?w=400&h=400&fit=crop', // Mouse
    3: 'https://images.unsplash.com/photo-1595044426077-d36d9236d54a?w=400&h=400&fit=crop', // Keyboard - Fixed
    4: 'https://images.unsplash.com/photo-1527443224154-c4a3942d3acf?w=400&h=400&fit=crop', // Monitor
    5: 'https://images.unsplash.com/photo-1505740420928-5e560c06d30e?w=400&h=400&fit=crop', // Headphones
    6: 'https://images.unsplash.com/photo-1587829741301-dc798b83Add3?w=400&h=400&fit=crop', // Webcam - Fixed
    7: 'https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?w=400&h=400&fit=crop', // Smartphone
    8: 'https://images.unsplash.com/photo-1561154464-82e9adf32764?w=400&h=400&fit=crop'  // Tablet
  };

  const getProductImage = (itemId) => {
    return productImages[itemId] || `https://images.unsplash.com/photo-1560472354-b33ff0c44a43?w=400&h=400&fit=crop`;
  };

  useEffect(() => {
    if (token) {
      const userData = localStorage.getItem('user');
      if (userData) {
        setUser(JSON.parse(userData));
      }
      setCurrentScreen('main');
      fetchItems();
      fetchCart();
      fetchOrders();
    }
  }, [token]);

  const fetchItems = async () => {
    try {
      setLoading(true);
      const response = await fetch(`${API_BASE_URL}/items`);
      const data = await response.json();
      console.log('Fetched items:', data);
      setItems(data || []);
    } catch (error) {
      console.error('Failed to fetch items:', error);
      // Enhanced fallback data with prices and descriptions
      setItems([
        { id: 1, name: 'MacBook Pro', status: 'active', price: 1299.99, description: 'Powerful laptop for professionals' },
        { id: 2, name: 'Wireless Mouse', status: 'active', price: 29.99, description: 'Ergonomic wireless mouse' },
        { id: 3, name: 'Mechanical Keyboard', status: 'active', price: 89.99, description: 'RGB mechanical gaming keyboard' },
        { id: 4, name: '4K Monitor', status: 'active', price: 299.99, description: '27-inch 4K UHD monitor' },
        { id: 5, name: 'Noise-Canceling Headphones', status: 'active', price: 199.99, description: 'Premium audio experience' },
        { id: 6, name: 'HD Webcam', status: 'active', price: 79.99, description: '1080p webcam with auto-focus' },
        { id: 7, name: 'iPhone 15 Pro', status: 'active', price: 999.99, description: 'Latest iPhone with A17 Pro chip' },
        { id: 8, name: 'iPad Air', status: 'active', price: 599.99, description: 'Lightweight tablet for creativity' }
      ]);
    } finally {
      setLoading(false);
    }
  };

  const fetchCart = async () => {
    if (!token) return;
    
    try {
      const response = await fetch(`${API_BASE_URL}/carts`, {
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      });

      if (response.ok) {
        const carts = await response.json();
        const activeCart = carts.find(cart => cart.status === 'active');
        setCart(activeCart || null);
        console.log('Fetched cart:', activeCart);
      }
    } catch (error) {
      console.error('Failed to fetch cart:', error);
    }
  };

  const fetchOrders = async () => {
    if (!token) return;
    
    try {
      const response = await fetch(`${API_BASE_URL}/orders`, {
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      });

      if (response.ok) {
        const ordersData = await response.json();
        setOrders(ordersData || []);
        console.log('Fetched orders:', ordersData);
      }
    } catch (error) {
      console.error('Failed to fetch orders:', error);
    }
  };

  const LoginScreen = () => {
    const [username, setUsername] = useState('testuser');
    const [password, setPassword] = useState('password');
    const [error, setError] = useState('');

    const handleLogin = async () => {
      setLoading(true);
      setError('');
      
      try {
        const response = await fetch(`${API_BASE_URL}/users/login`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ username, password }),
        });

        const data = await response.json();
        
        if (response.ok) {
          setToken(data.token);
          setUser({ id: data.user_id, username: data.username });
          localStorage.setItem('token', data.token);
          localStorage.setItem('user', JSON.stringify({ id: data.user_id, username: data.username }));
          setCurrentScreen('main');
          console.log('Login successful:', data);
        } else {
          setError(data.error || 'Login failed');
        }
      } catch (error) {
        console.error('Login error:', error);
        setError('Unable to connect to server. Please check if the backend is running on port 8080.');
      } finally {
        setLoading(false);
      }
    };

    const handleKeyPress = (e) => {
      if (e.key === 'Enter') {
        handleLogin();
      }
    };

    return (
      <div className="min-h-screen bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center">
        <div className="bg-white p-8 rounded-lg shadow-2xl w-96">
          <div className="text-center mb-8">
            <User className="mx-auto mb-4 text-blue-500" size={48} />
            <h1 className="text-2xl font-bold text-gray-800">Welcome Back</h1>
            <p className="text-gray-600">Sign in to your account</p>
          </div>
          
          {error && (
            <div className="mb-4 p-3 bg-red-100 border border-red-400 text-red-700 rounded">
              {error}
            </div>
          )}
          
          <div className="space-y-6">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">Username</label>
              <input
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                onKeyPress={handleKeyPress}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                disabled={loading}
                tabIndex={-1}
              />
            </div>
            
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">Password</label>
              <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                onKeyPress={handleKeyPress}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                disabled={loading}
              />
            </div>
            
            <button
              onClick={handleLogin}
              disabled={loading}
              className="w-full bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 disabled:opacity-50 transition-colors"
            >
              {loading ? 'Signing in...' : 'Sign In'}
            </button>
          </div>
          
          <div className="mt-4 text-sm text-gray-600 text-center">
            <p>Test credentials:</p>
            <p>Username: testuser | Password: password</p>
          </div>
        </div>
      </div>
    );
  };

  const Navigation = () => {
    const cartItemCount = cart?.items?.length || 0;

    return (
      <nav className="bg-white shadow-lg border-b">
        <div className="max-w-7xl mx-auto px-4">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center space-x-8">
              <div className="flex items-center space-x-3">
                <Package className="text-blue-500" size={28} />
                <h1 className="text-xl font-bold text-gray-800">ShopEasy</h1>
              </div>
              <div className="flex space-x-6">
                <button
                  onClick={() => setCurrentPage('home')}
                  className={`flex items-center space-x-1 px-3 py-2 rounded-md transition-colors ${
                    currentPage === 'home' ? 'text-blue-600 bg-blue-50' : 'text-gray-600 hover:text-gray-800'
                  }`}
                >
                  <Home size={18} />
                  <span>Home</span>
                </button>
                <button
                  onClick={() => setCurrentPage('cart')}
                  className={`flex items-center space-x-1 px-3 py-2 rounded-md transition-colors relative ${
                    currentPage === 'cart' ? 'text-blue-600 bg-blue-50' : 'text-gray-600 hover:text-gray-800'
                  }`}
                >
                  <ShoppingCart size={18} />
                  <span>Cart</span>
                  {cartItemCount > 0 && (
                    <span className="absolute -top-1 -right-1 bg-red-500 text-white text-xs rounded-full h-5 w-5 flex items-center justify-center">
                      {cartItemCount}
                    </span>
                  )}
                </button>
                <button
                  onClick={() => setCurrentPage('orders')}
                  className={`flex items-center space-x-1 px-3 py-2 rounded-md transition-colors ${
                    currentPage === 'orders' ? 'text-blue-600 bg-blue-50' : 'text-gray-600 hover:text-gray-800'
                  }`}
                >
                  <History size={18} />
                  <span>Orders</span>
                </button>
                <button
                  onClick={() => setCurrentPage('profile')}
                  className={`flex items-center space-x-1 px-3 py-2 rounded-md transition-colors ${
                    currentPage === 'profile' ? 'text-blue-600 bg-blue-50' : 'text-gray-600 hover:text-gray-800'
                  }`}
                >
                  <User size={18} />
                  <span>Profile</span>
                </button>
              </div>
            </div>
            <div className="flex items-center space-x-4">
              <span className="text-gray-600">Welcome, {user?.username}</span>
              <button
                onClick={logout}
                className="bg-red-500 text-white px-4 py-2 rounded-md hover:bg-red-600 transition-colors flex items-center space-x-2"
              >
                <LogOut size={18} />
                <span>Logout</span>
              </button>
            </div>
          </div>
        </div>
      </nav>
    );
  };

  const HomePage = () => {
    const [addingToCart, setAddingToCart] = useState(null);

    const addToCart = async (itemId) => {
      setAddingToCart(itemId);
      console.log('Token being sent:', token);
      console.log('Adding item to cart:', itemId);
      
      try {
        const response = await fetch(`${API_BASE_URL}/carts`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
          },
          body: JSON.stringify({ item_ids: [itemId] }),
        });

        console.log('Response status:', response.status);
        
        if (response.ok) {
          const updatedCart = await response.json();
          setCart(updatedCart);
          await fetchCart(); // Refresh cart data to get latest totals
          console.log('Item added to cart:', updatedCart);
        } else {
          const error = await response.json();
          console.error('Cart error response:', error);
          alert(`Failed to add item to cart: ${error.error || 'Unknown error'}`);
        }
      } catch (error) {
        console.error('Error adding item to cart:', error);
        alert('Error adding item to cart. Please check console for details.');
      } finally {
        setAddingToCart(null);
      }
    };

    return (
      <div className="max-w-7xl mx-auto px-4 py-8">
        <h2 className="text-3xl font-bold text-gray-800 mb-8">Featured Products</h2>
        
        {loading && items.length === 0 ? (
          <div className="text-center py-12">
            <Package className="mx-auto text-gray-400 mb-4 animate-pulse" size={64} />
            <p className="text-gray-600">Loading products...</p>
          </div>
        ) : (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
            {items.map(item => (
              <div key={item.id} className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow">
                <div className="aspect-square overflow-hidden">
                  <img
                    src={getProductImage(item.id)}
                    alt={item.name}
                    className="w-full h-full object-cover hover:scale-105 transition-transform duration-300"
                  />
                </div>
                <div className="p-4">
                  <h3 className="font-semibold text-lg text-gray-800 mb-2">{item.name}</h3>
                  <p className="text-gray-600 text-sm mb-3">{item.description || 'High quality product'}</p>
                  <div className="flex items-center mb-3">
                    <div className="flex items-center">
                      {[...Array(5)].map((_, i) => (
                        <Star
                          key={i}
                          size={16}
                          className={i < 4 ? 'text-yellow-400 fill-current' : 'text-gray-300'}
                        />
                      ))}
                    </div>
                    <span className="text-sm text-gray-600 ml-2">(4.0)</span>
                  </div>
                  <div className="flex items-center justify-between">
                    <span className="text-xl font-bold text-gray-800">
                      ${item.price || '99.99'}
                    </span>
                    <button
                      onClick={() => addToCart(item.id)}
                      disabled={addingToCart === item.id}
                      className="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors disabled:opacity-50 flex items-center space-x-2"
                    >
                      <ShoppingCart size={18} />
                      <span>
                        {addingToCart === item.id ? 'Adding...' : 'Add to Cart'}
                      </span>
                    </button>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    );
  };

  const CartPage = () => {
    const cartItems = cart?.items || [];
    
    // Debug: Log cart data to understand structure
    console.log('Cart data:', cart);
    console.log('Cart items:', cartItems);
    console.log('Available items for price lookup:', items);
    
    // Calculate total properly by finding the item details and multiplying by quantity
    const cartTotal = cartItems.reduce((total, cartItem) => {
      const item = items.find(item => item.id === cartItem.item_id);
      const price = item?.price || 0;
      const quantity = cartItem.quantity || 1; // Default to 1 if quantity not available
      console.log(`Item: ${item?.name}, Price: ${price}, Quantity: ${quantity}, Subtotal: ${price * quantity}`);
      return total + (price * quantity);
    }, 0);

    console.log('Cart total calculated:', cartTotal);

    const checkout = async () => {
      try {
        setLoading(true);
        
        if (!cart || !cart.items || cart.items.length === 0) {
          alert('No items in cart. Please add items to cart first.');
          return;
        }

        const orderResponse = await fetch(`${API_BASE_URL}/orders`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
          },
          body: JSON.stringify({ cart_id: cart.id }),
        });

        if (orderResponse.ok) {
          const order = await orderResponse.json();
          setCart(null);
          await fetchOrders(); // Refresh orders
          alert(`Order successful! 🎉\nOrder ID: ${order.id}\nItems: ${order.cart?.items?.length || 0}`);
          setCurrentPage('orders');
          console.log('Order created:', order);
        } else {
          const error = await orderResponse.json();
          alert(`Failed to create order: ${error.error || 'Unknown error'}`);
        }
      } catch (error) {
        console.error('Error during checkout:', error);
        alert('Error during checkout. Please check console for details.');
      } finally {
        setLoading(false);
      }
    };

    return (
      <div className="max-w-7xl mx-auto px-4 py-8">
        <h2 className="text-3xl font-bold text-gray-800 mb-8">Shopping Cart</h2>
        
        {cartItems.length === 0 ? (
          <div className="text-center py-12">
            <ShoppingCart size={64} className="mx-auto text-gray-400 mb-4" />
            <p className="text-gray-600 text-lg">Your cart is empty</p>
            <button
              onClick={() => setCurrentPage('home')}
              className="mt-4 bg-blue-600 text-white px-6 py-3 rounded-md hover:bg-blue-700"
            >
              Continue Shopping
            </button>
          </div>
        ) : (
          <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
            <div className="lg:col-span-2">
              {cartItems.map(cartItem => {
                const item = items.find(item => item.id === cartItem.item_id);
                const quantity = cartItem.quantity || 1;
                const itemTotal = (item?.price || 0) * quantity;
                
                return (
                  <div key={cartItem.id} className="bg-white rounded-lg shadow-md p-6 mb-4">
                    <div className="flex items-center space-x-4">
                      <img
                        src={getProductImage(cartItem.item_id)}
                        alt={item?.name || 'Product'}
                        className="w-20 h-20 object-cover rounded-md"
                      />
                      <div className="flex-1">
                        <h3 className="font-semibold text-lg text-gray-800">{item?.name || 'Unknown Item'}</h3>
                        <p className="text-gray-600">${item?.price || '0.00'} each</p>
                        <p className="text-sm text-gray-500">Quantity: {quantity}</p>
                        <p className="text-lg font-semibold text-blue-600">Total: ${itemTotal.toFixed(2)}</p>
                      </div>
                    </div>
                  </div>
                );
              })}
            </div>
            <div className="lg:col-span-1">
              <div className="bg-white rounded-lg shadow-md p-6 sticky top-4">
                <h3 className="font-semibold text-lg mb-4">Order Summary</h3>
                <div className="space-y-2 mb-4">
                  <div className="flex justify-between">
                    <span>Items ({cartItems.length})</span>
                    <span>${cartTotal.toFixed(2)}</span>
                  </div>
                  <div className="flex justify-between">
                    <span>Shipping</span>
                    <span>Free</span>
                  </div>
                  <div className="border-t pt-2 flex justify-between font-semibold">
                    <span>Total</span>
                    <span>${cartTotal.toFixed(2)}</span>
                  </div>
                </div>
                <button
                  onClick={checkout}
                  disabled={loading}
                  className="w-full bg-green-600 text-white py-3 rounded-md hover:bg-green-700 font-semibold disabled:opacity-50"
                >
                  {loading ? 'Processing...' : 'Place Order'}
                </button>
              </div>
            </div>
          </div>
        )}
      </div>
    );
  };

  const OrdersPage = () => (
    <div className="max-w-7xl mx-auto px-4 py-8">
      <h2 className="text-3xl font-bold text-gray-800 mb-8">Order History</h2>
      
      {orders.length === 0 ? (
        <div className="text-center py-12">
          <Package size={64} className="mx-auto text-gray-400 mb-4" />
          <p className="text-gray-600 text-lg">No orders yet</p>
          <button
            onClick={() => setCurrentPage('home')}
            className="mt-4 bg-blue-600 text-white px-6 py-3 rounded-md hover:bg-blue-700"
          >
            Start Shopping
          </button>
        </div>
      ) : (
        <div className="space-y-6">
          {orders.map(order => (
            <div key={order.id} className="bg-white rounded-lg shadow-md p-6">
              <div className="flex justify-between items-start mb-4">
                <div>
                  <h3 className="font-semibold text-lg">Order #{order.id}</h3>
                  <p className="text-gray-600">
                    Placed on {new Date(order.created_at).toLocaleDateString()}
                  </p>
                </div>
                <div className="text-right">
                  <span className="px-3 py-1 rounded-full text-sm font-medium bg-green-100 text-green-800">
                    Completed
                  </span>
                  <p className="text-lg font-semibold mt-2">
                    {order.cart?.items?.length || 0} items
                  </p>
                </div>
              </div>
              <div className="border-t pt-4">
                <h4 className="font-medium mb-2">Items:</h4>
                {order.cart?.items?.map((cartItem, index) => {
                  const item = items.find(item => item.id === cartItem.item_id);
                  return (
                    <div key={index} className="flex justify-between py-1">
                      <span>{item?.name || `Item ${cartItem.item_id}`}</span>
                      <span>${item?.price || '0.00'}</span>
                    </div>
                  );
                })}
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );

  const ProfilePage = () => (
    <div className="max-w-7xl mx-auto px-4 py-8">
      <h2 className="text-3xl font-bold text-gray-800 mb-8">My Profile</h2>
      
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <div className="lg:col-span-1">
          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="text-center">
              <div className="w-24 h-24 bg-gradient-to-r from-blue-500 to-purple-600 rounded-full mx-auto mb-4 flex items-center justify-center">
                <User size={40} className="text-white" />
              </div>
              <h3 className="text-xl font-semibold">{user?.username}</h3>
              <p className="text-gray-600">Premium Member</p>
            </div>
          </div>
        </div>
        
        <div className="lg:col-span-2">
          <div className="bg-white rounded-lg shadow-md p-6">
            <h3 className="text-xl font-semibold mb-6">Account Information</h3>
            <div className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Username</label>
                <input
                  type="text"
                  defaultValue={user?.username}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Email</label>
                <input
                  type="email"
                  defaultValue={`${user?.username}@example.com`}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Phone</label>
                <input
                  type="tel"
                  placeholder="+1 (555) 123-4567"
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Address</label>
                <textarea
                  placeholder="123 Main St, City, State 12345"
                  rows={3}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
              <button
                onClick={() => alert('Profile updated successfully!')}
                className="bg-blue-600 text-white px-6 py-2 rounded-md hover:bg-blue-700"
              >
                Update Profile
              </button>
            </div>
          </div>
          
          <div className="bg-white rounded-lg shadow-md p-6 mt-6">
            <h3 className="text-xl font-semibold mb-4">Quick Stats</h3>
            <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
              <div className="text-center p-4 bg-blue-50 rounded-lg">
                <Package size={24} className="mx-auto text-blue-600 mb-2" />
                <p className="text-2xl font-bold text-blue-600">{orders.length}</p>
                <p className="text-sm text-gray-600">Total Orders</p>
              </div>
              <div className="text-center p-4 bg-green-50 rounded-lg">
                <ShoppingCart size={24} className="mx-auto text-green-600 mb-2" />
                <p className="text-2xl font-bold text-green-600">{cart?.items?.length || 0}</p>
                <p className="text-sm text-gray-600">Cart Items</p>
              </div>
              <div className="text-center p-4 bg-purple-50 rounded-lg">
                <Star size={24} className="mx-auto text-purple-600 mb-2" />
                <p className="text-2xl font-bold text-purple-600">4.8</p>
                <p className="text-sm text-gray-600">Avg Rating</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );

  const logout = () => {
    setToken(null);
    setUser(null);
    setCart(null);
    setOrders([]);
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    setCurrentScreen('login');
    setCurrentPage('home');
  };

  const renderCurrentPage = () => {
    switch (currentPage) {
      case 'home':
        return <HomePage />;
      case 'cart':
        return <CartPage />;
      case 'orders':
        return <OrdersPage />;
      case 'profile':
        return <ProfilePage />;
      default:
        return <HomePage />;
    }
  };

  return (
    <div className="App">
      {currentScreen === 'login' && <LoginScreen />}
      {currentScreen === 'main' && (
        <div className="min-h-screen bg-gray-50">
          <Navigation />
          {renderCurrentPage()}
        </div>
      )}
    </div>
  );
};

export default App;