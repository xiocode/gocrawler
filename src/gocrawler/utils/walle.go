/**
 * Created with IntelliJ IDEA.
 * User: xio
 * Date: 13-2-19
 * Time: 下午4:32
 * To change this template use File | Settings | File Templates.
 */
package utils


func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
