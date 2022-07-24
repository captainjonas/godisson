package api

/**
* GMap uses serialized state of key instead of hashCode or equals methods.
* GMap doesn't allow to store null as key or value.
 */
type GMap interface {

	/**
	* Loads all map entries to this Redis map
	 */
	loadAll(replaceExistingValues bool, parallelism int32) error

	get(key interface{}) (interface{}, error)

	put(key, value interface{}) error

	putIfAbsent(key, value interface{}) error

	/**
	* Returns GLock instance associated with key
	 */
	getFairLock(key interface{}) (GLock, error)

	/**
	* Returns GLock instance associated with key
	 */
	getLock(key interface{}) (GLock, error)

	/**
	* Returns true if this map contains map entry
	 */
	containsKey(key interface{}) (bool, error)

	/**
	* Returns true if this map contains any map entry with specified value, otherwise false
	 */
	containsValue(value interface{}) (bool, error)

	/**
	* Removes map entry by specified key and returns value.
	 */
	remove(key interface{}) (interface{}, error)

	/**
	* @return previous associated value
	*           or null if there is no map entry stored before and doesn't store new map entry
	 */
	replace(key, value interface{}) (interface{}, error)

	/**
	* Replaces previous oldValue with a newValue mapped by specified key.
	* Returns false if previous value doesn't exist or equal to oldValue
	 */
	replaceWithNewValue(key, oldValue, newValue interface{}) (bool, error)
}
