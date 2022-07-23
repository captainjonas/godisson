package api

import "time"

type (
	GLock interface {

		//Returns name of object
		getName() (string, error)

		/**
		* Acquire the lock with difined leaseTime
		* Waits if necessary until lock becomes available
		*
		* Lock will be released automatically after defined leaseTime interval
		*
		*@param leaseTime the maximum time to hold the lock after it's acquisition,
		*        if it hasn't already been released by invoking <code>unlock</code>.
		*        If leaseTime is -1, hold the lock until explicitly unlocked.
		 */
		lockInterruptibly(leaseTime int64, unit time.Duration) error

		/**
		* Tries to acquire the lock with defined leaseTime.
		* Waits up to defined waitTime if necessary until the lock becomes available.
		* Lock will be released autmatically after defined leaseTime interval.
		 */
		tryLock(waitTime, leaseTime int64, unit time.Duration) (bool, error)

		lock(leaseTime int64, unit time.Duration) error

		/**
		* Unlocks the lock independently of its state
		 */
		forceUnlock() (bool, error)

		/**
		* Checks if the lock locked
		 */
		isLocked() (bool, error)

		/**
		* Remaining time to live of the lock
		* @return time in milliseconds
		*          -2 if the lock does not exist.
		*          -1 if the lock exists but has no associated expire.
		 */
		remainTimeToLive() (int64, error)
	}
)
