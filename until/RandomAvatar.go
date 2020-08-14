package until

import "github.com/123456/c_code"

func RandomAvatar() string {
	l := []string{"/_avatar/head-picture-1.jpg", "/_avatar/head-picture-10.jpg", "/_avatar/head-picture-100.jpg", "/_avatar/head-picture-101.jpg", "/_avatar/head-picture-102.jpg", "/_avatar/head-picture-103.jpg", "/_avatar/head-picture-104.jpg", "/_avatar/head-picture-105.jpg", "/_avatar/head-picture-106.jpg", "/_avatar/head-picture-107.jpg", "/_avatar/head-picture-108.jpg", "/_avatar/head-picture-109.jpg", "/_avatar/head-picture-11.jpg", "/_avatar/head-picture-110.jpg", "/_avatar/head-picture-111.jpg", "/_avatar/head-picture-112.jpg", "/_avatar/head-picture-113.jpg", "/_avatar/head-picture-114.jpg", "/_avatar/head-picture-115.jpg", "/_avatar/head-picture-116.jpg", "/_avatar/head-picture-117.jpg", "/_avatar/head-picture-118.jpg", "/_avatar/head-picture-119.jpg", "/_avatar/head-picture-12.jpg", "/_avatar/head-picture-120.jpg", "/_avatar/head-picture-121.jpg", "/_avatar/head-picture-122.jpg", "/_avatar/head-picture-123.jpg", "/_avatar/head-picture-124.jpg", "/_avatar/head-picture-125.jpg", "/_avatar/head-picture-126.jpg", "/_avatar/head-picture-127.jpg", "/_avatar/head-picture-128.jpg", "/_avatar/head-picture-129.jpg", "/_avatar/head-picture-13.jpg", "/_avatar/head-picture-130.jpg", "/_avatar/head-picture-131.jpg", "/_avatar/head-picture-132.jpg", "/_avatar/head-picture-133.jpg", "/_avatar/head-picture-134.jpg", "/_avatar/head-picture-135.jpg", "/_avatar/head-picture-136.jpg", "/_avatar/head-picture-137.jpg", "/_avatar/head-picture-138.jpg", "/_avatar/head-picture-139.jpg", "/_avatar/head-picture-14.jpg", "/_avatar/head-picture-140.jpg", "/_avatar/head-picture-141.jpg", "/_avatar/head-picture-142.jpg", "/_avatar/head-picture-143.jpg", "/_avatar/head-picture-144.jpg", "/_avatar/head-picture-145.jpg", "/_avatar/head-picture-146.jpg", "/_avatar/head-picture-147.jpg", "/_avatar/head-picture-148.jpg", "/_avatar/head-picture-149.jpg", "/_avatar/head-picture-15.jpg", "/_avatar/head-picture-150.jpg", "/_avatar/head-picture-151.jpg", "/_avatar/head-picture-152.jpg", "/_avatar/head-picture-153.jpg", "/_avatar/head-picture-154.jpg", "/_avatar/head-picture-155.jpg", "/_avatar/head-picture-156.jpg", "/_avatar/head-picture-157.jpg", "/_avatar/head-picture-158.jpg", "/_avatar/head-picture-159.jpg", "/_avatar/head-picture-16.jpg", "/_avatar/head-picture-160.jpg", "/_avatar/head-picture-161.jpg", "/_avatar/head-picture-162.jpg", "/_avatar/head-picture-163.jpg", "/_avatar/head-picture-164.jpg", "/_avatar/head-picture-165.jpg", "/_avatar/head-picture-166.jpg", "/_avatar/head-picture-167.jpg", "/_avatar/head-picture-168.jpg", "/_avatar/head-picture-169.jpg", "/_avatar/head-picture-17.jpg", "/_avatar/head-picture-170.jpg", "/_avatar/head-picture-171.jpg", "/_avatar/head-picture-172.jpg", "/_avatar/head-picture-173.jpg", "/_avatar/head-picture-174.jpg", "/_avatar/head-picture-175.jpg", "/_avatar/head-picture-176.jpg", "/_avatar/head-picture-177.jpg", "/_avatar/head-picture-178.jpg", "/_avatar/head-picture-179.jpg", "/_avatar/head-picture-18.jpg", "/_avatar/head-picture-180.jpg", "/_avatar/head-picture-181.jpg", "/_avatar/head-picture-182.jpg", "/_avatar/head-picture-183.jpg", "/_avatar/head-picture-184.jpg", "/_avatar/head-picture-185.jpg", "/_avatar/head-picture-186.jpg", "/_avatar/head-picture-187.jpg", "/_avatar/head-picture-188.jpg", "/_avatar/head-picture-189.jpg", "/_avatar/head-picture-19.jpg", "/_avatar/head-picture-190.jpg", "/_avatar/head-picture-191.jpg", "/_avatar/head-picture-192.jpg", "/_avatar/head-picture-193.jpg", "/_avatar/head-picture-194.jpg", "/_avatar/head-picture-195.jpg", "/_avatar/head-picture-196.jpg", "/_avatar/head-picture-197.jpg", "/_avatar/head-picture-198.jpg", "/_avatar/head-picture-199.jpg", "/_avatar/head-picture-2.jpg", "/_avatar/head-picture-20.jpg", "/_avatar/head-picture-200.jpg", "/_avatar/head-picture-201.jpg", "/_avatar/head-picture-202.jpg", "/_avatar/head-picture-203.jpg", "/_avatar/head-picture-204.jpg", "/_avatar/head-picture-205.jpg", "/_avatar/head-picture-206.jpg", "/_avatar/head-picture-207.jpg", "/_avatar/head-picture-208.jpg", "/_avatar/head-picture-209.jpg", "/_avatar/head-picture-21.jpg", "/_avatar/head-picture-210.jpg", "/_avatar/head-picture-211.jpg", "/_avatar/head-picture-212.jpg", "/_avatar/head-picture-213.jpg", "/_avatar/head-picture-214.jpg", "/_avatar/head-picture-215.jpg", "/_avatar/head-picture-216.jpg", "/_avatar/head-picture-217.jpg", "/_avatar/head-picture-218.jpg", "/_avatar/head-picture-219.jpg", "/_avatar/head-picture-22.jpg", "/_avatar/head-picture-220.jpg", "/_avatar/head-picture-221.jpg", "/_avatar/head-picture-222.jpg", "/_avatar/head-picture-223.jpg", "/_avatar/head-picture-224.jpg", "/_avatar/head-picture-225.jpg", "/_avatar/head-picture-226.jpg", "/_avatar/head-picture-227.jpg", "/_avatar/head-picture-228.jpg", "/_avatar/head-picture-229.jpg", "/_avatar/head-picture-23.jpg", "/_avatar/head-picture-230.jpg", "/_avatar/head-picture-231.jpg", "/_avatar/head-picture-232.jpg", "/_avatar/head-picture-233.jpg", "/_avatar/head-picture-234.jpg", "/_avatar/head-picture-235.jpg", "/_avatar/head-picture-236.jpg", "/_avatar/head-picture-237.jpg", "/_avatar/head-picture-238.jpg", "/_avatar/head-picture-239.jpg", "/_avatar/head-picture-24.jpg", "/_avatar/head-picture-240.jpg", "/_avatar/head-picture-241.jpg", "/_avatar/head-picture-242.jpg", "/_avatar/head-picture-243.jpg", "/_avatar/head-picture-244.jpg", "/_avatar/head-picture-245.jpg", "/_avatar/head-picture-246.jpg", "/_avatar/head-picture-247.jpg", "/_avatar/head-picture-248.jpg", "/_avatar/head-picture-249.jpg", "/_avatar/head-picture-25.jpg", "/_avatar/head-picture-250.jpg", "/_avatar/head-picture-251.jpg", "/_avatar/head-picture-252.jpg", "/_avatar/head-picture-253.jpg", "/_avatar/head-picture-254.jpg", "/_avatar/head-picture-255.jpg", "/_avatar/head-picture-256.jpg", "/_avatar/head-picture-257.jpg", "/_avatar/head-picture-258.jpg", "/_avatar/head-picture-259.jpg", "/_avatar/head-picture-26.jpg", "/_avatar/head-picture-260.jpg", "/_avatar/head-picture-261.jpg", "/_avatar/head-picture-262.jpg", "/_avatar/head-picture-263.jpg", "/_avatar/head-picture-264.jpg", "/_avatar/head-picture-265.jpg", "/_avatar/head-picture-266.jpg", "/_avatar/head-picture-267.jpg", "/_avatar/head-picture-268.jpg", "/_avatar/head-picture-269.jpg", "/_avatar/head-picture-27.jpg", "/_avatar/head-picture-270.jpg", "/_avatar/head-picture-271.jpg", "/_avatar/head-picture-272.jpg", "/_avatar/head-picture-273.jpg", "/_avatar/head-picture-274.jpg", "/_avatar/head-picture-275.jpg", "/_avatar/head-picture-276.jpg", "/_avatar/head-picture-277.jpg", "/_avatar/head-picture-278.jpg", "/_avatar/head-picture-279.jpg", "/_avatar/head-picture-28.jpg", "/_avatar/head-picture-280.jpg", "/_avatar/head-picture-281.jpg", "/_avatar/head-picture-282.jpg", "/_avatar/head-picture-283.jpg", "/_avatar/head-picture-284.jpg", "/_avatar/head-picture-285.jpg", "/_avatar/head-picture-286.jpg", "/_avatar/head-picture-287.jpg", "/_avatar/head-picture-288.jpg", "/_avatar/head-picture-289.jpg", "/_avatar/head-picture-29.jpg", "/_avatar/head-picture-290.jpg", "/_avatar/head-picture-291.jpg", "/_avatar/head-picture-292.jpg", "/_avatar/head-picture-293.jpg", "/_avatar/head-picture-294.jpg", "/_avatar/head-picture-295.jpg", "/_avatar/head-picture-296.jpg", "/_avatar/head-picture-297.jpg", "/_avatar/head-picture-298.jpg", "/_avatar/head-picture-299.jpg", "/_avatar/head-picture-3.jpg", "/_avatar/head-picture-30.jpg", "/_avatar/head-picture-300.jpg", "/_avatar/head-picture-301.jpg", "/_avatar/head-picture-302.jpg", "/_avatar/head-picture-303.jpg", "/_avatar/head-picture-304.jpg", "/_avatar/head-picture-305.jpg", "/_avatar/head-picture-306.jpg", "/_avatar/head-picture-307.jpg", "/_avatar/head-picture-308.jpg", "/_avatar/head-picture-309.jpg", "/_avatar/head-picture-31.jpg", "/_avatar/head-picture-310.jpg", "/_avatar/head-picture-311.jpg", "/_avatar/head-picture-312.jpg", "/_avatar/head-picture-313.jpg", "/_avatar/head-picture-314.jpg", "/_avatar/head-picture-315.jpg", "/_avatar/head-picture-316.jpg", "/_avatar/head-picture-317.jpg", "/_avatar/head-picture-318.jpg", "/_avatar/head-picture-319.jpg", "/_avatar/head-picture-32.jpg", "/_avatar/head-picture-320.jpg", "/_avatar/head-picture-321.jpg", "/_avatar/head-picture-322.jpg", "/_avatar/head-picture-323.jpg", "/_avatar/head-picture-324.jpg", "/_avatar/head-picture-325.jpg", "/_avatar/head-picture-326.jpg", "/_avatar/head-picture-327.jpg", "/_avatar/head-picture-328.jpg", "/_avatar/head-picture-329.jpg", "/_avatar/head-picture-33.jpg", "/_avatar/head-picture-330.jpg", "/_avatar/head-picture-331.jpg", "/_avatar/head-picture-332.jpg", "/_avatar/head-picture-333.jpg", "/_avatar/head-picture-334.jpg", "/_avatar/head-picture-335.jpg", "/_avatar/head-picture-336.jpg", "/_avatar/head-picture-337.jpg", "/_avatar/head-picture-338.jpg", "/_avatar/head-picture-339.jpg", "/_avatar/head-picture-34.jpg", "/_avatar/head-picture-340.jpg", "/_avatar/head-picture-341.jpg", "/_avatar/head-picture-342.jpg", "/_avatar/head-picture-343.jpg", "/_avatar/head-picture-344.jpg", "/_avatar/head-picture-345.jpg", "/_avatar/head-picture-346.jpg", "/_avatar/head-picture-347.jpg", "/_avatar/head-picture-348.jpg", "/_avatar/head-picture-349.jpg", "/_avatar/head-picture-35.jpg", "/_avatar/head-picture-350.jpg", "/_avatar/head-picture-351.jpg", "/_avatar/head-picture-352.jpg", "/_avatar/head-picture-353.jpg", "/_avatar/head-picture-354.jpg", "/_avatar/head-picture-355.jpg", "/_avatar/head-picture-356.jpg", "/_avatar/head-picture-357.jpg", "/_avatar/head-picture-358.jpg", "/_avatar/head-picture-359.jpg", "/_avatar/head-picture-36.jpg", "/_avatar/head-picture-360.jpg", "/_avatar/head-picture-361.jpg", "/_avatar/head-picture-362.jpg", "/_avatar/head-picture-363.jpg", "/_avatar/head-picture-364.jpg", "/_avatar/head-picture-365.jpg", "/_avatar/head-picture-366.jpg", "/_avatar/head-picture-367.jpg", "/_avatar/head-picture-368.jpg", "/_avatar/head-picture-369.jpg", "/_avatar/head-picture-37.jpg", "/_avatar/head-picture-370.jpg", "/_avatar/head-picture-371.jpg", "/_avatar/head-picture-372.jpg", "/_avatar/head-picture-373.jpg", "/_avatar/head-picture-374.jpg", "/_avatar/head-picture-375.jpg", "/_avatar/head-picture-376.jpg", "/_avatar/head-picture-377.jpg", "/_avatar/head-picture-378.jpg", "/_avatar/head-picture-379.jpg", "/_avatar/head-picture-38.jpg", "/_avatar/head-picture-380.jpg", "/_avatar/head-picture-381.jpg", "/_avatar/head-picture-382.jpg", "/_avatar/head-picture-383.jpg", "/_avatar/head-picture-384.jpg", "/_avatar/head-picture-385.jpg", "/_avatar/head-picture-386.jpg", "/_avatar/head-picture-387.jpg", "/_avatar/head-picture-388.jpg", "/_avatar/head-picture-389.jpg", "/_avatar/head-picture-39.jpg", "/_avatar/head-picture-390.jpg", "/_avatar/head-picture-391.jpg", "/_avatar/head-picture-392.jpg", "/_avatar/head-picture-393.jpg", "/_avatar/head-picture-394.jpg", "/_avatar/head-picture-395.jpg", "/_avatar/head-picture-396.jpg", "/_avatar/head-picture-397.jpg", "/_avatar/head-picture-398.jpg", "/_avatar/head-picture-399.jpg", "/_avatar/head-picture-4.jpg", "/_avatar/head-picture-40.jpg", "/_avatar/head-picture-400.jpg", "/_avatar/head-picture-401.jpg", "/_avatar/head-picture-402.jpg", "/_avatar/head-picture-403.jpg", "/_avatar/head-picture-404.jpg", "/_avatar/head-picture-405.jpg", "/_avatar/head-picture-406.jpg", "/_avatar/head-picture-407.jpg", "/_avatar/head-picture-408.jpg", "/_avatar/head-picture-409.jpg", "/_avatar/head-picture-41.jpg", "/_avatar/head-picture-410.jpg", "/_avatar/head-picture-411.jpg", "/_avatar/head-picture-412.jpg", "/_avatar/head-picture-413.jpg", "/_avatar/head-picture-414.jpg", "/_avatar/head-picture-415.jpg", "/_avatar/head-picture-416.jpg", "/_avatar/head-picture-417.jpg", "/_avatar/head-picture-418.jpg", "/_avatar/head-picture-419.jpg", "/_avatar/head-picture-42.jpg", "/_avatar/head-picture-420.jpg", "/_avatar/head-picture-421.jpg", "/_avatar/head-picture-422.jpg", "/_avatar/head-picture-423.jpg", "/_avatar/head-picture-424.jpg", "/_avatar/head-picture-425.jpg", "/_avatar/head-picture-426.jpg", "/_avatar/head-picture-427.jpg", "/_avatar/head-picture-428.jpg", "/_avatar/head-picture-429.jpg", "/_avatar/head-picture-43.jpg", "/_avatar/head-picture-430.jpg", "/_avatar/head-picture-431.jpg", "/_avatar/head-picture-432.jpg", "/_avatar/head-picture-433.jpg", "/_avatar/head-picture-434.jpg", "/_avatar/head-picture-435.jpg", "/_avatar/head-picture-436.jpg", "/_avatar/head-picture-437.jpg", "/_avatar/head-picture-438.jpg", "/_avatar/head-picture-439.jpg", "/_avatar/head-picture-44.jpg", "/_avatar/head-picture-440.jpg", "/_avatar/head-picture-441.jpg", "/_avatar/head-picture-442.jpg", "/_avatar/head-picture-443.jpg", "/_avatar/head-picture-444.jpg", "/_avatar/head-picture-445.jpg", "/_avatar/head-picture-446.jpg", "/_avatar/head-picture-447.jpg", "/_avatar/head-picture-448.jpg", "/_avatar/head-picture-449.jpg", "/_avatar/head-picture-45.jpg", "/_avatar/head-picture-450.jpg", "/_avatar/head-picture-451.jpg", "/_avatar/head-picture-452.jpg", "/_avatar/head-picture-453.jpg", "/_avatar/head-picture-454.jpg", "/_avatar/head-picture-455.jpg", "/_avatar/head-picture-456.jpg", "/_avatar/head-picture-457.jpg", "/_avatar/head-picture-458.jpg", "/_avatar/head-picture-459.jpg", "/_avatar/head-picture-46.jpg", "/_avatar/head-picture-460.jpg", "/_avatar/head-picture-461.jpg", "/_avatar/head-picture-462.jpg", "/_avatar/head-picture-463.jpg", "/_avatar/head-picture-464.jpg", "/_avatar/head-picture-465.jpg", "/_avatar/head-picture-466.jpg", "/_avatar/head-picture-467.jpg", "/_avatar/head-picture-468.jpg", "/_avatar/head-picture-469.jpg", "/_avatar/head-picture-47.jpg", "/_avatar/head-picture-470.jpg", "/_avatar/head-picture-471.jpg", "/_avatar/head-picture-472.jpg", "/_avatar/head-picture-473.jpg", "/_avatar/head-picture-474.jpg", "/_avatar/head-picture-475.jpg", "/_avatar/head-picture-476.jpg", "/_avatar/head-picture-477.jpg", "/_avatar/head-picture-478.jpg", "/_avatar/head-picture-479.jpg", "/_avatar/head-picture-48.jpg", "/_avatar/head-picture-480.jpg", "/_avatar/head-picture-481.jpg", "/_avatar/head-picture-482.jpg", "/_avatar/head-picture-483.jpg", "/_avatar/head-picture-484.jpg", "/_avatar/head-picture-485.jpg", "/_avatar/head-picture-486.jpg", "/_avatar/head-picture-487.jpg", "/_avatar/head-picture-488.jpg", "/_avatar/head-picture-489.jpg", "/_avatar/head-picture-49.jpg", "/_avatar/head-picture-490.jpg", "/_avatar/head-picture-491.jpg", "/_avatar/head-picture-492.jpg", "/_avatar/head-picture-493.jpg", "/_avatar/head-picture-494.jpg", "/_avatar/head-picture-495.jpg", "/_avatar/head-picture-496.jpg", "/_avatar/head-picture-497.jpg", "/_avatar/head-picture-498.jpg", "/_avatar/head-picture-499.jpg", "/_avatar/head-picture-5.jpg", "/_avatar/head-picture-50.jpg", "/_avatar/head-picture-500.jpg", "/_avatar/head-picture-501.jpg", "/_avatar/head-picture-502.jpg", "/_avatar/head-picture-503.jpg", "/_avatar/head-picture-504.jpg", "/_avatar/head-picture-505.jpg", "/_avatar/head-picture-506.jpg", "/_avatar/head-picture-507.jpg", "/_avatar/head-picture-508.jpg", "/_avatar/head-picture-509.jpg", "/_avatar/head-picture-51.jpg", "/_avatar/head-picture-510.jpg", "/_avatar/head-picture-511.jpg", "/_avatar/head-picture-512.jpg", "/_avatar/head-picture-513.jpg", "/_avatar/head-picture-514.jpg", "/_avatar/head-picture-515.jpg", "/_avatar/head-picture-516.jpg", "/_avatar/head-picture-517.jpg", "/_avatar/head-picture-518.jpg", "/_avatar/head-picture-519.jpg", "/_avatar/head-picture-52.jpg", "/_avatar/head-picture-520.jpg", "/_avatar/head-picture-521.jpg", "/_avatar/head-picture-522.jpg", "/_avatar/head-picture-523.jpg", "/_avatar/head-picture-524.jpg", "/_avatar/head-picture-525.jpg", "/_avatar/head-picture-526.jpg", "/_avatar/head-picture-527.jpg", "/_avatar/head-picture-528.jpg", "/_avatar/head-picture-529.jpg", "/_avatar/head-picture-53.jpg", "/_avatar/head-picture-530.jpg", "/_avatar/head-picture-531.jpg", "/_avatar/head-picture-532.jpg", "/_avatar/head-picture-533.jpg", "/_avatar/head-picture-534.jpg", "/_avatar/head-picture-535.jpg", "/_avatar/head-picture-536.jpg", "/_avatar/head-picture-537.jpg", "/_avatar/head-picture-538.jpg", "/_avatar/head-picture-539.jpg", "/_avatar/head-picture-54.jpg", "/_avatar/head-picture-540.jpg", "/_avatar/head-picture-541.jpg", "/_avatar/head-picture-542.jpg", "/_avatar/head-picture-543.jpg", "/_avatar/head-picture-544.jpg", "/_avatar/head-picture-545.jpg", "/_avatar/head-picture-546.jpg", "/_avatar/head-picture-547.jpg", "/_avatar/head-picture-548.jpg", "/_avatar/head-picture-549.jpg", "/_avatar/head-picture-55.jpg", "/_avatar/head-picture-550.jpg", "/_avatar/head-picture-551.jpg", "/_avatar/head-picture-552.jpg", "/_avatar/head-picture-553.jpg", "/_avatar/head-picture-554.jpg", "/_avatar/head-picture-555.jpg", "/_avatar/head-picture-556.jpg", "/_avatar/head-picture-557.jpg", "/_avatar/head-picture-558.jpg", "/_avatar/head-picture-559.jpg", "/_avatar/head-picture-56.jpg", "/_avatar/head-picture-560.jpg", "/_avatar/head-picture-561.jpg", "/_avatar/head-picture-562.jpg", "/_avatar/head-picture-563.jpg", "/_avatar/head-picture-564.jpg", "/_avatar/head-picture-565.jpg", "/_avatar/head-picture-566.jpg", "/_avatar/head-picture-567.jpg", "/_avatar/head-picture-568.jpg", "/_avatar/head-picture-569.jpg", "/_avatar/head-picture-57.jpg", "/_avatar/head-picture-570.jpg", "/_avatar/head-picture-571.jpg", "/_avatar/head-picture-572.jpg", "/_avatar/head-picture-573.jpg", "/_avatar/head-picture-574.jpg", "/_avatar/head-picture-575.jpg", "/_avatar/head-picture-576.jpg", "/_avatar/head-picture-577.jpg", "/_avatar/head-picture-578.jpg", "/_avatar/head-picture-579.jpg", "/_avatar/head-picture-58.jpg", "/_avatar/head-picture-580.jpg", "/_avatar/head-picture-581.jpg", "/_avatar/head-picture-582.jpg", "/_avatar/head-picture-583.jpg", "/_avatar/head-picture-584.jpg", "/_avatar/head-picture-585.jpg", "/_avatar/head-picture-586.jpg", "/_avatar/head-picture-587.jpg", "/_avatar/head-picture-588.jpg", "/_avatar/head-picture-589.jpg", "/_avatar/head-picture-59.jpg", "/_avatar/head-picture-590.jpg", "/_avatar/head-picture-591.jpg", "/_avatar/head-picture-592.jpg", "/_avatar/head-picture-593.jpg", "/_avatar/head-picture-594.jpg", "/_avatar/head-picture-595.jpg", "/_avatar/head-picture-596.jpg", "/_avatar/head-picture-597.jpg", "/_avatar/head-picture-598.jpg", "/_avatar/head-picture-599.jpg", "/_avatar/head-picture-6.jpg", "/_avatar/head-picture-60.jpg", "/_avatar/head-picture-600.jpg", "/_avatar/head-picture-601.jpg", "/_avatar/head-picture-602.jpg", "/_avatar/head-picture-604.jpg", "/_avatar/head-picture-605.jpg", "/_avatar/head-picture-606.jpg", "/_avatar/head-picture-607.jpg", "/_avatar/head-picture-608.jpg", "/_avatar/head-picture-61.jpg", "/_avatar/head-picture-610.jpg", "/_avatar/head-picture-611.jpg", "/_avatar/head-picture-612.jpg", "/_avatar/head-picture-613.jpg", "/_avatar/head-picture-614.jpg", "/_avatar/head-picture-615.jpg", "/_avatar/head-picture-616.jpg", "/_avatar/head-picture-617.jpg", "/_avatar/head-picture-618.jpg", "/_avatar/head-picture-619.jpg", "/_avatar/head-picture-62.jpg", "/_avatar/head-picture-620.jpg", "/_avatar/head-picture-621.jpg", "/_avatar/head-picture-622.jpg", "/_avatar/head-picture-623.jpg", "/_avatar/head-picture-624.jpg", "/_avatar/head-picture-625.jpg", "/_avatar/head-picture-626.jpg", "/_avatar/head-picture-627.jpg", "/_avatar/head-picture-628.jpg", "/_avatar/head-picture-629.jpg", "/_avatar/head-picture-63.jpg", "/_avatar/head-picture-630.jpg", "/_avatar/head-picture-631.jpg", "/_avatar/head-picture-632.jpg", "/_avatar/head-picture-633.jpg", "/_avatar/head-picture-634.jpg", "/_avatar/head-picture-635.jpg", "/_avatar/head-picture-636.jpg", "/_avatar/head-picture-637.jpg", "/_avatar/head-picture-638.jpg", "/_avatar/head-picture-639.jpg", "/_avatar/head-picture-64.jpg", "/_avatar/head-picture-640.jpg", "/_avatar/head-picture-641.jpg", "/_avatar/head-picture-642.jpg", "/_avatar/head-picture-65.jpg", "/_avatar/head-picture-66.jpg", "/_avatar/head-picture-67.jpg", "/_avatar/head-picture-68.jpg", "/_avatar/head-picture-69.jpg", "/_avatar/head-picture-7.jpg", "/_avatar/head-picture-70.jpg", "/_avatar/head-picture-71.jpg", "/_avatar/head-picture-72.jpg", "/_avatar/head-picture-73.jpg", "/_avatar/head-picture-74.jpg", "/_avatar/head-picture-75.jpg", "/_avatar/head-picture-76.jpg", "/_avatar/head-picture-77.jpg", "/_avatar/head-picture-78.jpg", "/_avatar/head-picture-79.jpg", "/_avatar/head-picture-8.jpg", "/_avatar/head-picture-80.jpg", "/_avatar/head-picture-81.jpg", "/_avatar/head-picture-82.jpg", "/_avatar/head-picture-83.jpg", "/_avatar/head-picture-84.jpg", "/_avatar/head-picture-85.jpg", "/_avatar/head-picture-86.jpg", "/_avatar/head-picture-87.jpg", "/_avatar/head-picture-88.jpg", "/_avatar/head-picture-89.jpg", "/_avatar/head-picture-9.jpg", "/_avatar/head-picture-90.jpg", "/_avatar/head-picture-91.jpg", "/_avatar/head-picture-92.jpg", "/_avatar/head-picture-93.jpg", "/_avatar/head-picture-94.jpg", "/_avatar/head-picture-95.jpg", "/_avatar/head-picture-96.jpg", "/_avatar/head-picture-97.jpg", "/_avatar/head-picture-98.jpg", "/_avatar/head-picture-99.jpg"}
	return l[c_code.Rand(len(l))]
}