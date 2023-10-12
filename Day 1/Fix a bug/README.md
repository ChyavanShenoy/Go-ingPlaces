# Problem: Fix a bug

## Problem Statement

Users are reporting that we're billing them for wildly inaccurate aounts. They're supposed to be billed for `$0.02` for each text message sent.

## Solution

The problem is that we're using the wrong operator to multiply the number of text messages by the cost per text message. Instead of using `*`, we're using `+`. This means that we're adding the number of text messages to the cost per text message, instead of multiplying them.

To fix this, we need to change the `+` to a `*` on line 18.

You can see the difference in the 2 files:
error_code.go and fixed_code.go
