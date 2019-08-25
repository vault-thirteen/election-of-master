////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

//+build test

package eom

import (
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_checkServiceName(t *testing.T) {

	var aTest *tester.Test
	var err error
	var serviceName string

	aTest = tester.New(t)

	// Test #1. Good Name.
	serviceName = "abcXYZ_123"
	err = checkServiceName(serviceName)
	aTest.MustBeNoError(err)

	// Test #2. Bad Name
	serviceName = "ы"
	err = checkServiceName(serviceName)
	aTest.MustBeAnError(err)

	// Test #3. Bad Name
	serviceName = "%"
	err = checkServiceName(serviceName)
	aTest.MustBeAnError(err)
}
